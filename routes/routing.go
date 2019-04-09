package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"todo-go/models"

	"github.com/gin-gonic/gin"
)

func GetTasks(c *gin.Context) {
	fmt.Println("success to get tasks")
	c.JSON(http.StatusOK, models.GetTasks())
}

func DeleteTask(c *gin.Context) {
	id, err := strconv.Atoi(c.PostForm("id"))
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, models.DeleteTask(id))
}

func AddTask(c *gin.Context) {
	taskName := c.PostForm("task_name")
	c.JSON(http.StatusOK, models.AddTask(taskName))
}
