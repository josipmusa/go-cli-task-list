package cmd

import (
	"encoding/json"
	"fmt"
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
		data, err := json.MarshalIndent(note, "", "  ")
		if err != nil {
			return err
		}
		fmt.Println(string(data))
		return nil
	},
}

func init() {
	ViewNoteCmd.Flags().IntVar(&id, "id", 0, "Note ID")
	_ = ViewNoteCmd.MarkFlagRequired("id")
}
