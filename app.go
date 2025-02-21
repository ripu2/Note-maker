package main

import (
	"fmt"

	"example.com/note-taker/packages/notes"
)

func main() {

	for {
		if !notes.RunNotes() {
			break
		}
	}
	fmt.Println("Exiting the program...")
}
