package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"go-task-manager/internal/storage"
	"go-task-manager/internal/task"
)

type Config struct {
	TaskFile string
}

func NewConfig() Config {
	return Config{
		TaskFile: "tasks.json",
	}
}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println(`  go run . add "Task title"`)
	fmt.Println("  go run . list")
	fmt.Println("  go run . complete <id>")
	fmt.Println("  go run . delete <id>")
}

func handleAdd(service *task.Service) {
	if len(os.Args) < 3 {
		fmt.Println("Task title is required")
		return
	}

	title := strings.Join(os.Args[2:], " ")

	newTask, err := service.Add(title)
	if err != nil {
		if errors.Is(err, task.ErrEmptyTitle) {
			fmt.Println(err)
			return
		}

		fmt.Println("Error adding task:", err)
		return
	}

	fmt.Printf("Task %d added successfully\n", newTask.ID)
}

func handleList(service *task.Service) {
	tasks, err := service.List()
	if err != nil {
		fmt.Println("Error listing tasks:", err)
		return
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found")
		return
	}

	for _, currentTask := range tasks {
		status := "[ ]"

		if currentTask.Completed {
			status = "[x]"
		}

		fmt.Printf(
			"%s %d - %s\n",
			status,
			currentTask.ID,
			currentTask.Title,
		)
	}
}

func handleComplete(service *task.Service) {
	id, ok := parseIDArgument()
	if !ok {
		return
	}

	if err := service.Complete(id); err != nil {
		var notFoundError task.NotFoundError
		if errors.As(err, &notFoundError) {
			fmt.Printf("Task %d not found\n", notFoundError.ID)
			return
		}

		fmt.Println("Error completing task:", err)
		return
	}

	fmt.Printf("Task %d completed\n", id)
}

func handleDelete(service *task.Service) {
	id, ok := parseIDArgument()
	if !ok {
		return
	}

	if err := service.Delete(id); err != nil {
		var notFoundError task.NotFoundError
		if errors.As(err, &notFoundError) {
			fmt.Printf("Task %d not found\n", notFoundError.ID)
			return
		}

		fmt.Println("Error deleting task:", err)
		return
	}

	fmt.Printf("Task %d deleted\n", id)
}

func parseIDArgument() (int, bool) {
	if len(os.Args) < 3 {
		fmt.Println("Task ID is required")
		return 0, false
	}

	id, err := strconv.Atoi(os.Args[2])
	if err != nil || id <= 0 {
		fmt.Println("Task ID must be a positive number")
		return 0, false
	}

	return id, true
}

func main() {
	cfg := NewConfig()

	if value := os.Getenv("TASKS_FILE"); value != "" {
		cfg.TaskFile = value
	}

	repository := storage.NewJSONRepository(cfg.TaskFile)
	service := task.NewService(repository)

	if len(os.Args) < 2 {
		printUsage()
		return
	}

	switch os.Args[1] {
	case "add":
		handleAdd(service)
	case "list":
		handleList(service)
	case "complete":
		handleComplete(service)
	case "delete":
		handleDelete(service)
	default:
		fmt.Println("Unknown command:", os.Args[1])
		printUsage()
	}
}
