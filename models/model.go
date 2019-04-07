package models

import (
	"fmt"

	"github.com/kshiva1126/todo-go/manipulations"
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
	t.Id = id
	t.Name = name
	return
}

func GetTasks() (tc TaskCollection) {
	tc = TaskCollection{}
	tasks := manipulations.DbGetAll()
	for _, v := range tasks {
		task := doTask(v.Id, v.Name)
		tc.Tasks = append(tc.Tasks, &task)
	}
	fmt.Println(tc.Tasks)
	return
}
