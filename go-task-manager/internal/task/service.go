package task

import (
	"errors"
	"fmt"
	"strings"
)

type NotFoundError struct {
	ID int
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("task %d not found", e.ID)
}

var ErrEmptyTitle = errors.New("task title cannot be empty")

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (service *Service) Add(title string) (Task, error) {
	title = strings.TrimSpace(title)

	if title == "" {
		return Task{}, ErrEmptyTitle
	}

	tasks, err := service.repository.Load()
	if err != nil {
		return Task{}, err
	}

	newTask := New(NextID(tasks), title)
	tasks = append(tasks, newTask)

	if err := service.repository.Save(tasks); err != nil {
		return Task{}, err
	}

	return newTask, nil
}

func (service *Service) List() ([]Task, error) {
	return service.repository.Load()
}

func NextID(tasks []Task) int {
	maxID := 0

	for _, currentTask := range tasks {
		if currentTask.ID > maxID {
			maxID = currentTask.ID
		}
	}

	return maxID + 1
}

func (service *Service) Complete(id int) error {
	tasks, err := service.repository.Load()
	if err != nil {
		return err
	}

	if !CompleteByID(tasks, id) {
		return NotFoundError{ID: id}
	}

	return service.repository.Save(tasks)
}

func (service *Service) Delete(id int) error {
	tasks, err := service.repository.Load()
	if err != nil {
		return err
	}

	updatedTasks, found := DeleteByID(tasks, id)
	if !found {
		return NotFoundError{ID: id}
	}

	return service.repository.Save(updatedTasks)
}

func CompleteByID(tasks []Task, id int) bool {
	for index := range tasks {
		if tasks[index].ID == id {
			tasks[index].Complete()
			return true
		}
	}

	return false
}

func DeleteByID(tasks []Task, id int) ([]Task, bool) {
	for index, currentTask := range tasks {
		if currentTask.ID == id {
			return append(
				tasks[:index],
				tasks[index+1:]...,
			), true
		}
	}

	return tasks, false
}
