/*
Copyright © 2024 Oleksandr Voronkov <dev.ovoronkov@gmail.com>
*/
package cmd

import (
	"dev-oleksandrv/todogo-cli/internal/config"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "todogo",
	Short: "todogo is Command Line Application for Managing To-Dos",
}

func Execute() {
	config.LoadConfig()

	listCmd.AddCommand(listsLsCmd)
	listCmd.AddCommand(listsCreateCmd)
	listCmd.AddCommand(listsCheckoutCmd)
	listCmd.AddCommand(listsRemoveCmd)

	rootCmd.AddCommand(tasksLsCmd)

	rootCmd.AddCommand(listCmd)

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


