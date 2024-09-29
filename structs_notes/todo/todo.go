package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

// pointer isnt used because it isnt editing/updating the note
// pointer is use to avoid creating a copy and updating the copy
func (todo Todo) Display() {
	fmt.Println(todo.Text)
}

func (todo Todo) Save() error {
	fileName := "todo.json"
	json, err := json.Marshal(todo)

	if err != nil {
		return err
	}

	// owner read and write permission, others read only
	// there's a possiblity of writing file failurre
	return os.WriteFile(fileName, json, 0644)
}

func New(content string) (Todo, error) {
	if content == "" {
		return Todo{}, errors.New("Invalid input")
	}

	return Todo{
		Text: content,
	}, nil
}
