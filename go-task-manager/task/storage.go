package task

import (
	"bytes"
	"encoding/json"
	"os"
)

func Save(filename string, tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0064)
}

func Load(filename string) ([]Task, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil
		}

		return nil, err
	}

	if len(bytes.TrimSpace(data)) == 0 {
		return []Task{}, nil
	}

	var tasks []Task

	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}
