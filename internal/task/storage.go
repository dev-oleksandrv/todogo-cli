package task

import (
	"encoding/csv"
	"log"
	"os"
)

const (
	taskStorageFilename = "tasks.csv"
)

func CreateTaskStorageFile() {
	if _, err := os.Stat(taskStorageFilename); os.IsNotExist(err) {
		file, err := os.Create(taskStorageFilename)
		if err != nil {
			log.Fatalf(`Cannot setup storage for tasks: %s`, err)
		}
		defer file.Close()

		writer := csv.NewWriter(file)
		defer writer.Flush()
		
		if err := writer.Write([]string{"ID", "ListID", "Content", "Completed", "UpdatedAt"}); err != nil {
			log.Fatalf("Cannot configure storage for tasks: %s", err)
		}
	}
}
