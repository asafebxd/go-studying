package task

type Task struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func New(id int, title string) Task {
	return Task{
		ID:        id,
		Title:     title,
		Completed: false,
	}
}

func (task *Task) Complete() {
	task.Completed = true
}

func NextID(tasks []Task) int {
	id := 0

	for _, task := range tasks {
		if task.ID > id {
			id = task.ID
		}
	}

	return id + 1
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
			updatedTasks := append(
				tasks[:index],
				tasks[index+1:]...,
			)

			return updatedTasks, true
		}
	}
	return tasks, false
}
