package cmd

import (
	"fmt"
	"go-cli-task-list/services"

	"github.com/spf13/cobra"
)

var DeleteNoteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Remove a note by ID",
	RunE: func(cmd *cobra.Command, args []string) error {
		noteService := services.NewNoteService()
		err := noteService.DeleteNote(id)
		if err != nil {
			return err
		}
		fmt.Println("Note removed successfully")
		return nil
	},
}

func init() {
	DeleteNoteCmd.Flags().IntVar(&id, "id", 0, "Note ID")
	_ = DeleteNoteCmd.MarkFlagRequired("id")
}
