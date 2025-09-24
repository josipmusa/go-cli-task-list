package services

import (
	"encoding/json"
	"fmt"
	"go-cli-task-list/internal/models"
	"os"
	"time"
)

type NoteService struct{}

const fileName = "notes.json"

func NewNoteService() *NoteService {
	return &NoteService{}
}

func (service *NoteService) SaveNote(note models.Note) error {
	notes, err := service.LoadNotes()
	if err != nil {
		return err
	}
	note.CreatedAt = time.Now().UTC()
	id := 0
	if len(notes) > 0 {
		id = notes[len(notes)-1].ID + 1
	}
	note.ID = id
	notes = append(notes, note)
	return service.writeToFile(notes)
}

func (service *NoteService) LoadNotes() ([]models.Note, error) {
	var notes []models.Note
	data, err := os.ReadFile(fileName)
	if err != nil {
		return []models.Note{}, err
	}

	if len(data) != 0 {
		err = json.Unmarshal(data, &notes)
		if err != nil {
			return []models.Note{}, err
		}
	}
	return notes, nil
}

func (service *NoteService) FindNote(id int) (models.Note, error) {
	notes, err := service.LoadNotes()
	if err != nil {
		return models.Note{}, err
	}
	if len(notes) == 0 {
		return models.Note{}, fmt.Errorf("note with id %d not found", id)
	}
	for _, note := range notes {
		if note.ID == id {
			return note, nil
		}
	}
	return models.Note{}, fmt.Errorf("note with id %d not found", id)
}

func (service *NoteService) DeleteNote(id int) error {
	notes, err := service.LoadNotes()
	if err != nil {
		return err
	}

	for i, note := range notes {
		if note.ID == id {
			notes = append(notes[:i], notes[i+1:]...)
			return service.writeToFile(notes)
		}
	}
	return fmt.Errorf("note with id %d not found", id)
}

func (service *NoteService) writeToFile(notes []models.Note) error {
	newData, err := json.MarshalIndent(notes, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, newData, 0644)
}
