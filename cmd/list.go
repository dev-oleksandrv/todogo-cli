/*
Copyright © 2024 Oleksandr Voronkov <dev.ovoronkov@gmail.com>
*/
package cmd

import (
	"dev-oleksandrv/todogo-cli/internal/datetime"
	"dev-oleksandrv/todogo-cli/internal/list"
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"

	"github.com/golang-module/carbon/v2"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Manage todo lists",
}

var listsLsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all todo lists",
	Run: func(cmd *cobra.Command, args []string) {
		lists := list.GetLists()
		if len(lists) == 0 {
			fmt.Println("You have no created lists yet. Use todogo list create command to create one")
			return
		}
		writer := tabwriter.NewWriter(os.Stdout, 10, 0, 2, ' ', 0)
		fmt.Fprintln(writer, "ID\tName\tCreated At")
		for _, list := range lists {
			fmt.Fprintf(writer, "%d\t%s\t%s\n", list.ID, list.Name, carbon.Parse(datetime.GetFormattedTime(list.CreatedAt)).DiffForHumans())
		}
		writer.Flush()
	},
}

var listsCurrentCmd = &cobra.Command{
	Use:   "current",
	Short: "Returns a current list",
	Run: func(cmd *cobra.Command, args []string) {
		currentList := list.GetCurrentList()
		if currentList == nil {
			fmt.Println("You have no current list")
			return
		}
		fmt.Printf("You are on \"%s\" list with ID %d", currentList.Name, currentList.ID)
	},
}

var listsCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new list",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		listName := args[0]
		id := list.CreateList(listName)
		fmt.Printf("A new list '%s' was successfully created with ID: %d\n", listName, id)
		if err := list.CheckoutList(id); err != nil {
			fmt.Printf("Cannot checkout to a list with ID: %d\n", id)
			return
		}
		fmt.Printf("You were checkout to list with ID %d\n", id)
	},
}

var listsCheckoutCmd = &cobra.Command{
	Use:   "checkout",
	Short: "Checkout to a list by ID",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Argument is incorrect")
			return
		}
		if err := list.CheckoutList(id); err != nil {
			fmt.Printf("Cannot checkout to a list with ID: %d\n", id)
			return
		}
		fmt.Printf("You were checkout to list with ID %d\n", id)
	},
}

var listsRemoveCmd = &cobra.Command{
	Use: "rm",
	Short: "Removes a list by ID",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Argument is incorrect")
			return
		}
		if err := list.RemoveList(id); err != nil {
			fmt.Printf("Cannot remove a list with ID: %d\n", id)
			return
		}
		fmt.Printf("You have successfully deleted list with %d\n", id)
	},
}
