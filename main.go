package main

import (
	"go-cli-task-list/cmd"
	"go-cli-task-list/services"
)

func main() {
	services.CreateFile()
	_ = cmd.RootCmd.Execute()
}
