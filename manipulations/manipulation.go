package manipulations

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Task struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func GormConnect() *gorm.DB {
	DBMS := "mysql"
	USER := "root"
	PASS := "root"
	PROTOCOL := "tcp(127.0.0.1:3306)"
	DBNAME := "todo_db"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err)
	}
	return db
}

func DbGetAll() []Task {
	var tasks []Task
	db := GormConnect()
	db.Find(&tasks)
	db.LogMode(true)
	db.Close()
	return tasks
}
