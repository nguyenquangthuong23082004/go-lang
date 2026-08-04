package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

func main() {
	title, content := getNoteData()
	todoText := getUserInput("Todo text:")

	todo, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	todo.Display()
	err = todo.Save()
	if err != nil {
		fmt.Println("Saving the todo failed:", err)
		return
	}
	fmt.Println("Saving the todo succeeded!")

	userNote.Display()
	err = userNote.Save()
	if err != nil {
		fmt.Println("Saving the note failed:", err)
		return
	}
	fmt.Println("Saving the note succeeded!")
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")
	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}

	// Cắt bỏ ký tự xuống dòng ở cuối chuỗi
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
