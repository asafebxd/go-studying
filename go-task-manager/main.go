package main

import (
	"fmt"
	"go-task-manager/task"
	"os"
	"strconv"
	"strings"
)

const filename = "tasks.json"

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println(`  go run . add "Task title"`)
	fmt.Println("  go run . list")
	fmt.Println("  go run . complete <id>")
	fmt.Println("  go run . delete <id>")
}

func handleAdd() {
	if len(os.Args) < 3 {
		fmt.Println("Task title is required")
		return
	}

	title := strings.TrimSpace(strings.Join(os.Args[2:], " "))
	if title == "" {
		fmt.Println("Task title cannot be empty")
		return
	}

	tasks, err := task.Load(filename)
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	newTask := task.New(task.NextID(tasks), title)
	tasks = append(tasks, newTask)

	if err := task.Save(filename, tasks); err != nil {
		fmt.Println("Error saving tasks", err)
		return
	}

	fmt.Printf("Task %d added successfully\n", newTask.ID)
}

func handleList() {
	tasks, err := task.Load(filename)
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found")
		return
	}

	for _, currenctTask := range tasks {
		status := "[ ]"

		if currenctTask.Completed {
			status = "[x]"
		}

		fmt.Printf(
			"%s %d - %s\n",
			status,
			currenctTask.ID,
			currenctTask.Title,
		)
	}

}

func handleComplete() {
	id, ok := parseIDArgument()
	if !ok {
		return
	}

	if len(os.Args) < 3 {
		fmt.Println("Task ID is required")
		return
	}

	id, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Task ID must be a number")
		return
	}

	tasks, err := task.Load(filename)
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	found := task.CompleteByID(tasks, id)
	if !found {
		fmt.Printf("Task %d not found\n", id)
		return
	}

	if err := task.Save(filename, tasks); err != nil {
		fmt.Println("Erro saving tasks:", err)
		return
	}

	fmt.Printf("Task %d completed\n", id)
}

func handleDelete() {
	id, ok := parseIDArgument()
	if !ok {
		return
	}
	
	if len(os.Args) < 3 {
		fmt.Println("Task ID is required")
		return
	}

	id, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Task ID must be a number:", err)
		return
	}

	tasks, err := task.Load(filename)
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	tasks, found := task.DeleteByID(tasks, id)
	if !found {
		fmt.Printf("Task %d not found\n", id)
		return
	}

	if err := task.Save(filename, tasks); err != nil {
		fmt.Println("Error saving tasks:", err)
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
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		handleAdd()
	case "list":
		handleList()
	case "complete":
		handleComplete()
	case "delete":
		handleDelete()
	default:
		fmt.Println("Unkown command:", command)
		printUsage()
	}

}
