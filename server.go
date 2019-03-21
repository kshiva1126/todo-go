package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kshiva1126/todo-go/routes"
)

func main() {
	router := gin.Default()
	router.StaticFile("/", "templates/index.html")
	router.GET("/tasks", routes.GetTasks)
	router.Run(":8083")
}
