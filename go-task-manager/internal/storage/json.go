package storage

import (
	"bytes"
	"encoding/json"
	"os"

	"go-task-manager/internal/task"
)

type JSONRepository struct {
	filename string
}

func NewJSONRepository(filename string) *JSONRepository {
	return &JSONRepository{
		filename: filename,
	}
}

func (repository *JSONRepository) Load() ([]task.Task, error) {
	data, err := os.ReadFile(repository.filename)
	if err != nil {
		if os.IsNotExist(err) {
			return []task.Task{}, nil
		}

		return nil, err
	}

	if len(bytes.TrimSpace(data)) == 0 {
		return []task.Task{}, nil
	}

	var tasks []task.Task

	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (repository *JSONRepository) Save(tasks []task.Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(repository.filename, data, 0644)
}
