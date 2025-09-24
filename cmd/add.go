package cmd

import (
	"fmt"
	"go-cli-task-list/models"
	"go-cli-task-list/services"

	"github.com/spf13/cobra"
)

var (
	title string
	body  string
	tags  []string
)

var AddNoteCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a note",
	RunE: func(cmd *cobra.Command, args []string) error {
		note := models.Note{
			Title: title,
			Body:  body,
			Tags:  tags,
		}

		noteService := services.NewNoteService()
		err := noteService.SaveNote(note)
		if err != nil {
			return err
		}

		fmt.Println("Note added successfully")
		return nil
	},
}

func init() {
	AddNoteCmd.Flags().StringVar(&title, "title", "", "Title of the note")
	AddNoteCmd.Flags().StringVar(&body, "body", "", "Body of the note")
	AddNoteCmd.Flags().StringSliceVar(&tags, "tags", []string{}, "Tags of the note")
	_ = AddNoteCmd.MarkFlagRequired("title")
	_ = AddNoteCmd.MarkFlagRequired("body")
	_ = AddNoteCmd.MarkFlagRequired("tags")
}
