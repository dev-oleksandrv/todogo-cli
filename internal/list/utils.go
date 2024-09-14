package list

import (
	"dev-oleksandrv/todogo-cli/internal/config"
	"encoding/csv"
	"errors"
	"log"
	"os"
	"strings"
)

func GetCurrentList() *List {
	currentList := config.GetCurrentList()
	if currentList == -1 {
		return nil
	}

	lists := GetLists()
	if len(lists) == 0 {
		return nil
	}

	var list List
	for _, l := range lists {
		if l.ID == currentList {
			list = l
		}
	}

	return &list
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
	list := NewList(id, name)
	if err := writer.Write(list.ToRecord()); err != nil {
		log.Fatalf("Failed to create a list: %s", err)
	}
	config.IncrementListId()
	return id
}

func RemoveList(id int) error {
	lists, rtr := GetLists(), -1
	for index, list := range lists {
		if list.ID == id {
			rtr = index
			break
		}
	}
	if rtr == -1 {
		return errors.New("cannot find list with such id")
	}
	lists = append(lists[:rtr], lists[rtr+1:]...)
	file, err := os.Create(listStorageFilename)
	if err != nil {
		return err
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	records := [][]string{}
	for _, list := range lists {
		records = append(records, list.ToRecord())
	}
	err = writer.WriteAll(records)
	if err != nil {
		return err
	}
	writer.Flush()
	currentList := config.GetCurrentList()
	if currentList == id {
		if len(lists) > 0 {
			config.CheckoutList(lists[0].ID)
		} else {
			config.CheckoutList(-1)
		}
	}
	return nil
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
	if len(records) == 0 {
		return result
	}
	for _, record := range records[1:] {
		l, err := FromRecord(record)
		if err != nil {
			log.Fatalf("Cannot read record from storage: %s", err)
		}
		result = append(result, *l)
	}
	return result
}