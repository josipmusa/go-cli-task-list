package cmd

import (
	"encoding/json"
	"fmt"
	"go-cli-task-list/internal/models"
	"go-cli-task-list/internal/services"

	"github.com/spf13/cobra"
)

var (
	title string
	body  string
	tags  []string
	id    int
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
			for _, note := range notes {
				data, err := json.MarshalIndent(note, "", "  ")
				if err != nil {
					return err
				}
				fmt.Println(string(data))
			}
		}
		return nil
	},
}

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
	AddNoteCmd.Flags().StringVar(&title, "title", "", "Title of the note")
	AddNoteCmd.Flags().StringVar(&body, "body", "", "Body of the note")
	AddNoteCmd.Flags().StringSliceVar(&tags, "tags", []string{}, "Tags of the note")
	_ = AddNoteCmd.MarkFlagRequired("title")
	_ = AddNoteCmd.MarkFlagRequired("body")
	_ = AddNoteCmd.MarkFlagRequired("tags")

	ViewNoteCmd.Flags().IntVar(&id, "id", 0, "Note ID")
	_ = ViewNoteCmd.MarkFlagRequired("id")

	DeleteNoteCmd.Flags().IntVar(&id, "id", 0, "Note ID")
	_ = DeleteNoteCmd.MarkFlagRequired("id")
}
