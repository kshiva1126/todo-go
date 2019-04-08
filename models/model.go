package models

import (
	"todo-go/manipulations"
)

type Task struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type TaskCollection struct {
	Tasks []*Task `json:"items"`
}

func doTask(id int, name string) (t Task) {
	t = Task{}
	t.Id, t.Name = id, name
	return
}

func GetTasks() (tc TaskCollection) {
	tc = TaskCollection{}
	tasks := manipulations.GetAllTasks()
	for _, v := range tasks {
		task := doTask(v.Id, v.Name)
		tc.Tasks = append(tc.Tasks, &task)
	}
	return
}

func DeleteTask(id int) (tc TaskCollection) {
	manipulations.DeleteTask(id)
	tc = TaskCollection{}
	tc = GetTasks()
	return
}
