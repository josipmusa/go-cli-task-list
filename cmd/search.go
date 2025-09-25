package cmd

import (
	"encoding/json"
	"fmt"
	"go-cli-task-list/services"

	"github.com/spf13/cobra"
)

var SearchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for notes",
	Args:  cobra.ExactArgs(1), // enforce exactly 1 argument
	RunE: func(cmd *cobra.Command, args []string) error {
		noteService := services.NewNoteService()
		keyword := args[0]
		fmt.Println("Searching for:", keyword)
		notes, err := noteService.SearchNote(keyword)
		if err != nil {
			return err
		}
		if len(notes) == 0 {
			fmt.Println("No notes found")
		}

		for _, note := range notes {
			data, err := json.MarshalIndent(note, "", "  ")
			if err != nil {
				return err
			}
			fmt.Println(string(data))
		}

		return nil
	},
}
