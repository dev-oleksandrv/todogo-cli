package task

import (
	"dev-oleksandrv/todogo-cli/internal/config"
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Task struct {
	ID        int
	ListID int
	Content      string
	Completed bool
	UpdatedAt time.Time
}

func GetTasksInList() []Task {
	result := []Task{}
	tasks := GetTasks()
	currentListId := config.GetCurrentList()
	for _, task := range tasks {
		if task.ListID == currentListId {
			result = append(result, task)
		}
	}
	return result
}

func GetTasks() []Task {
	data, err := os.ReadFile(taskStorageFilename)
	if err != nil {
		log.Fatalf("Cannot open lists storage file: %s", err)
	}
	reader := csv.NewReader(strings.NewReader(string(data)))
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Cannot open lists storage file: %s", err)
	}
	result := []Task{}
	for _, record := range records[1:] {
		// parsedTime, err := time.Parse(time.RFC3339, record[2])
		// if err != nil {
		// 	log.Fatalf("Cannot parse CreatedAt: %s", err)
		// }
		parsedID, err := strconv.Atoi(record[0])
		if err != nil {
			log.Fatalf("Cannot parse ID: %s", err)
		}
		result = append(result, Task{
			ID: parsedID,
			ListID: 0,
			Content: record[2],
			Completed: false,
			UpdatedAt: time.Now(),
		})
	}
	return result
}
