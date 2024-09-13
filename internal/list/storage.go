package list

import (
	"encoding/csv"
	"log"
	"os"
)

const (
	listStorageFilename = "lists.csv"
)

func CreateListStorageFile() {
	if _, err := os.Stat(listStorageFilename); os.IsNotExist(err) {
		file, err := os.Create(listStorageFilename)
		if err != nil {
			log.Fatalf(`Cannot setup storage for lists: %s`, err)
		}
		defer file.Close()

		writer := csv.NewWriter(file)
		defer writer.Flush()
		
		if err := writer.Write([]string{"ID", "Name", "CreatedAt"}); err != nil {
			log.Fatalf("Cannot configure storage for list: %s", err)
		}
	}
}
