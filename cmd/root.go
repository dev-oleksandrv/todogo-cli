/*
Copyright © 2024 Oleksandr Voronkov <dev.ovoronkov@gmail.com>
*/
package cmd

import (
	"dev-oleksandrv/todogo-cli/internal/config"
	"dev-oleksandrv/todogo-cli/internal/list"
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"

	"github.com/golang-module/carbon/v2"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "todogo",
	Short: "todogo is Command Line Application for Managing To-Dos",
}

var listCmd = &cobra.Command{
	Use: "list",
	Short: "Manage todo lists",
}

var listsLsCmd = &cobra.Command{
	Use: "ls",
	Short: "List all todo lists",
	Run: func(cmd *cobra.Command, args []string) {
		writer := tabwriter.NewWriter(os.Stdout, 10, 0, 2, ' ', 0)
		lists := list.GetLists()

		fmt.Fprintln(writer, "ID\tName\tCreated At")
		for _, list := range lists {
			fmt.Fprintf(writer, "%d\t%s\t%s\n", list.ID, list.Name, carbon.Parse("2020-08-05 13:14:15").DiffForHumans())
		}
		writer.Flush()
	},
}

var listsCreateCmd = &cobra.Command{
	Use: "create",
	Short: "Create a new list",
	Args: cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		listName := args[0]
		id := list.CreateList(listName)
		fmt.Printf("A new list '%s' was successfully created with ID: %d\n", listName, id)
		if err := list.CheckoutList(id); err != nil {
			fmt.Printf("Cannot checkout to a list with ID: %d\n", id)
		}
		fmt.Printf("You were checkout to list with ID %d\n", id)
	},
}

var listsCheckoutCmd = &cobra.Command{
	Use: "checkout",
	Short: "Checkout to a list by ID",
	Args: cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Argument is incorrect")
		}
		if err := list.CheckoutList(id); err != nil {
			fmt.Printf("Cannot checkout to a list with ID: %d\n", id)
		}
		fmt.Printf("You were checkout to list with ID %d\n", id)
	},
}

func Execute() {
	config.LoadConfig()

	listCmd.AddCommand(listsLsCmd)
	listCmd.AddCommand(listsCreateCmd)
	listCmd.AddCommand(listsCheckoutCmd)

	rootCmd.AddCommand(listCmd)

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.todogo-cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


