package cmd

import (
	"go-cli-task-list/models"
	"go-cli-task-list/services"

	"github.com/spf13/cobra"
)

var ViewNoteCmd = &cobra.Command{
	Use:   "view",
	Short: "Show note details by ID",
	RunE: func(cmd *cobra.Command, args []string) error {
		noteService := services.NewNoteService()
		note, err := noteService.FindNote(id)
		if err != nil {
			return err
		}

		err = noteService.PrintNotes([]models.Note{note})
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	ViewNoteCmd.Flags().IntVar(&id, "id", 0, "Note ID")
	_ = ViewNoteCmd.MarkFlagRequired("id")
}
