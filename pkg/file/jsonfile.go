package file

import (
	"encoding/json"
	"io"
	"os"

	"github.com/pliniogsnascimento/little-habits/pkg/habit"
)

type JsonFileOpts struct {
	Path     string
	Filename string
}

func GetData(filepath string) (*habit.Habit, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	b, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var habit habit.Habit
	json.Unmarshal(b, habit)

	return &habit, nil
}
