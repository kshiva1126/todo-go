package main

import (
	"todo-go/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.StaticFile("/", "templates/index.html")
	router.GET("/tasks", routes.GetTasks)
	router.POST("/delete", routes.DeleteTask)
	router.Run(":8083")
}
