/*
Copyright © 2024 Oleksandr Voronkov <dev.ovoronkov@gmail.com>
*/
package cmd

import (
	"dev-oleksandrv/todogo-cli/internal/task"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/golang-module/carbon/v2"
	"github.com/spf13/cobra"
)

var tasksLsCmd = &cobra.Command{
	Use:   "ls",
	Short: "Display all incompleted todos in current list",
	Run: func(cmd *cobra.Command, args []string) {
		tasks := task.GetTasksInList()
		writer := tabwriter.NewWriter(os.Stdout, 10, 0, 2, ' ', 0)

		fmt.Fprintln(writer, "ID\tName\tUpdated At")
		for _, task := range tasks {
			fmt.Fprintf(writer, "%d\t%s\t%s\n", task.ID, task.Content, carbon.Parse("2020-08-05 13:14:15").DiffForHumans())
		}
		writer.Flush()
	},
}

var tasksAddCmd = &cobra.Command{
	Use: "add",
	Aliases: []string{"a"},
	Short: "Creates a task in a current list",
	Run: func(cmd *cobra.Command, args []string) {
		tasks := task.GetTasksInList()
		writer := tabwriter.NewWriter(os.Stdout, 10, 0, 2, ' ', 0)

		fmt.Fprintln(writer, "ID\tName\tUpdated At")
		for _, task := range tasks {
			fmt.Fprintf(writer, "%d\t%s\t%s\n", task.ID, task.Content, carbon.Parse("2020-08-05 13:14:15").DiffForHumans())
		}
		writer.Flush()
	},
}