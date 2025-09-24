package services

import (
	"fmt"
	"os"
)

func CreateFile() {
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		file, err := os.Create(fileName)
		if err != nil {
			fmt.Println("Error creating file", err)
			return
		}
		defer file.Close()
	}
}
