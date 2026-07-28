package storage

import (
	"bytes"
	"encoding/json"
	"fmt"
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

		return nil, fmt.Errorf("read tasks file: %w", err)
	}

	if len(bytes.TrimSpace(data)) == 0 {
		return []task.Task{}, nil
	}

	var tasks []task.Task

	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, fmt.Errorf("unmarshal tasks file: %w", err)
	}

	return tasks, nil
}

func (repository *JSONRepository) Save(tasks []task.Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal tasks: %w", err)
	}

	if err := os.WriteFile(repository.filename, data, 0644); err != nil {
		return fmt.Errorf("write tasks file: %w", err)
	}

	return nil
}
