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
