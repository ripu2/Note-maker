package communication

import (
	"errors"
	"fmt"
)

func TakeUserInput(prompt string) (string, error) {
	fmt.Printf("%s: ", prompt)
	var value string
	fmt.Scanln(&value)
	if value == "" {
		return "", errors.New("please provide a valid input")
	}
	return value, nil
}

func PrintNotesIOMenu() (choice string, err error) {
	fmt.Println("\nChoose an action:")
	fmt.Println("1. Create a new note")
	fmt.Println("2. Read a note")
	fmt.Println("3. Update a note")
	fmt.Println("4. Delete a note")
	fmt.Println("5. Quit")
	choice, _ = TakeUserInput("Enter Your Choice")
	return
}
