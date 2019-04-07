package manipulations

import (
	"fmt"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Task struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func GormConnect() *gorm.DB {
	connectionString := getConnectionString()
	db, err := gorm.Open("mysql", connectionString)

	if err != nil {
		panic(err)
	}
	return db
}

func getConnectionString() string {
	host := getParamString("MYSQL_DB_HOST", "localhost")
	port := getParamString("MYSQL_PORT", "3306")
	user := getParamString("MYSQL_USER", "root")
	pass := getParamString("MYSQL_PASSWORD", "root")
	dbname := getParamString("MYSQL_DB", "todo_db")
	protocol := getParamString("MYSQL_PROTOCOL", "tcp")
	dbargs := getParamString("MYSQL_DBARGS", " ")

	if strings.Trim(dbargs, " ") != "" {
		dbargs = "?" + dbargs
	} else {
		dbargs = ""
	}
	return fmt.Sprintf("%s:%s@%s([%s]:%s)/%s%s", user, pass, protocol, host, port, dbname, dbargs)
}

func getParamString(param string, defaultValue string) string {
	env := os.Getenv(param)
	if env != "" {
		return env
	}
	return defaultValue
}

func DbGetAll() []Task {
	var tasks []Task
	db := GormConnect()
	db.Find(&tasks)
	db.LogMode(true)
	db.Close()
	return tasks
}
