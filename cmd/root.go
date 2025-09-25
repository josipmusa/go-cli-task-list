package cmd

import "github.com/spf13/cobra"

var id int

var RootCmd = &cobra.Command{
	Use:   "pkb",
	Short: "Personal Knowledge Base stores notes",
}

func init() {
	RootCmd.AddCommand(AddNoteCmd, ReadNotesCmd, ViewNoteCmd, DeleteNoteCmd, SearchCmd)
}
