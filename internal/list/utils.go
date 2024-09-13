package list

import (
	"dev-oleksandrv/todogo-cli/internal/config"
	"encoding/csv"
	"errors"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type List struct {
	ID        int
	Name      string
	CreatedAt time.Time
}

func CreateList(name string) int {
	file, err := os.OpenFile(listStorageFilename, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Fatalf("Cannot open lists storage file: %s", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	id := config.GetListId()
	record := []string{strconv.Itoa(id), name, time.Now().String()}
	if err := writer.Write(record); err != nil {
		log.Fatalf("Failed to create a list: %s", err)
	}
	config.IncrementListId()
	return id
}

func CheckoutList(id int) error {
	lists := GetLists()
	for _, list := range lists {
		if list.ID == id {
			config.CheckoutList(id)
			return nil
		}
	}
	return errors.New("cannot find list with such id")
}

func GetLists() []List {
	data, err := os.ReadFile(listStorageFilename)
	if err != nil {
		log.Fatalf("Cannot open lists storage file: %s", err)
	}
	reader := csv.NewReader(strings.NewReader(string(data)))
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Cannot open lists storage file: %s", err)
	}
	result := []List{}
	for _, record := range records[1:] {
		// parsedTime, err := time.Parse(time.RFC3339, record[2])
		// if err != nil {
		// 	log.Fatalf("Cannot parse CreatedAt: %s", err)
		// }
		parsedID, err := strconv.Atoi(record[0])
		if err != nil {
			log.Fatalf("Cannot parse ID: %s", err)
		}
		result = append(result, List{
			ID: parsedID,
			Name: record[1],
			CreatedAt: time.Now(),
		})
	}
	return result
}