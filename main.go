/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"dev-oleksandrv/todogo-cli/cmd"
	"dev-oleksandrv/todogo-cli/internal/list"
	"dev-oleksandrv/todogo-cli/internal/task"
)

func main() {
	list.CreateListStorageFile()
	task.CreateTaskStorageFile()

	cmd.Execute()
}
