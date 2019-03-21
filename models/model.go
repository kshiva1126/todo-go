package models

type Task struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type TaskCollection struct {
	Tasks []Task `json:"items"`
}

func GetTasks() (tc TaskCollection) {
	tc = TaskCollection{
		[]Task{
			{1, "洗濯"},
			{2, "掃除"},
		},
	}

	return
}
