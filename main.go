package main

import (
	"go-cli-task-list/internal/cmd"
	"go-cli-task-list/internal/services"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pkb",
	Short: "Personal Knowledge Base stores notes",
}

func main() {
	services.CreateFile()
	rootCmd.AddCommand(cmd.AddNoteCmd)
	rootCmd.AddCommand(cmd.ReadNotesCmd)
	rootCmd.AddCommand(cmd.ViewNoteCmd)
	rootCmd.AddCommand(cmd.DeleteNoteCmd)

	_ = rootCmd.Execute()
}
