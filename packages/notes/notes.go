package notes

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"example.com/note-taker/packages/communication"
	"example.com/note-taker/packages/fileops"
)

type NotesType struct {
	Title     string
	Content   string
	CreatedAt time.Time
}

var notesFileName = "Notes.json"

func noteConstructor(title string, content string) (*NotesType, error) {
	if title == "" || content == "" {
		return nil, errors.New("title and content must not be empty")
	}
	return &NotesType{
		title,
		content,
		time.Now(),
	}, nil
}

func takeNotesInput() (title, contnent string) {
	noteTitle, _ := communication.TakeUserInput("Please Enter the tile of the note")
	noteContent, _ := communication.TakeUserInput("Please Enter the content of the note")
	return noteTitle, noteContent
}

func getNotes() (*NotesType, error) {
	fileStream, err := fileops.ReadFromFile(notesFileName)
	if err != nil {
		return nil, errors.New("you fucked up dude, something went wrong while reading from file")
	}
	decodedNoteValue, err := fileops.JSONDecoder[NotesType](fileStream)
	if err != nil {
		panic(err)
	}

	return decodedNoteValue, nil
}
func updateNotes(id string) (bool, error) {
	noteTitle, _ := communication.TakeUserInput("Please Enter the updated title of the note")
	noteContent, _ := communication.TakeUserInput("Please Enter the updated content of the note")
	currentNote, _ := getNotes()
	currentNote.Title = noteTitle
	currentNote.Content = noteContent

	jsonData, err := fileops.JSONEncoder(currentNote)
	if err != nil {
		errors.New("Failed to encode JSON data")
		return false, err
	}
	error := fileops.WriteToFile(notesFileName, jsonData)
	if error != nil {
		errors.New("Failed to write to file")
		return false, error
	}
	return true, nil
}

func delteNotes() error {
	if err := fileops.WriteToFile(notesFileName, []byte("{}")); err != nil {
		errors.New("Failed to write to file")
		return err
	}
	fmt.Println("All notes have been deleted.")
	return nil
}

func RunNotes() bool {
	var input, _ = communication.PrintNotesIOMenu()
	var choice, _ = strconv.ParseFloat(input, 64)

	switch choice {
	case 1:
		noteTitle, noteContent := takeNotesInput()
		var newNote *NotesType
		newNote, err := noteConstructor(noteTitle, noteContent)
		if err != nil {
			panic(err)
		} else {
			fmt.Printf("Your note has been successfully created with the following details:\nTitle: %s\nContent: %s\nCreated At: %s\n", newNote.Title, newNote.Content, newNote.CreatedAt)
			jsonData, _ := fileops.JSONEncoder(newNote)
			err := fileops.WriteToFile(notesFileName, jsonData)
			if err != nil {
				panic(err)
			}
		}
		return true
	case 2:
		val, err := getNotes()
		if err != nil {
			panic(err)
		}
		fmt.Printf("Title: %s\nContent: %s\nCreated At: %s\n", val.Title, val.Content, val.CreatedAt)
		return true
	case 3:
		id, _ := communication.TakeUserInput("Enter Notes Id")
		_, err := updateNotes(id)
		if err != nil {
			panic(err)
		} else {
			fmt.Println("Updated Successfully !!!")
		}
		return true
	case 4:
		err := delteNotes()
		if err != nil {
			panic(err)
		}
		return true
	default:
		return false
	}
}
