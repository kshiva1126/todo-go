package routes

import (
	"net/http"

	"todo-go/models"

	"github.com/gin-gonic/gin"
)

func GetTasks(c *gin.Context) {
	c.JSON(http.StatusOK, models.GetTasks())
}
