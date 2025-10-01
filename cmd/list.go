package cmd

import (
	"fmt"
	"go-cli-task-list/services"

	"github.com/spf13/cobra"
)

var ReadNotesCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all notes",
	RunE: func(cmd *cobra.Command, args []string) error {
		noteService := services.NewNoteService()
		notes, err := noteService.LoadNotes()
		if err != nil {
			return err
		}

		if len(notes) == 0 {
			fmt.Println("No notes found")
		} else {
			fmt.Println("Notes:")
			err := noteService.PrintNotes(notes)
			if err != nil {
				return err
			}
		}
		return nil
	},
}
