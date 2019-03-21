package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kshiva1126/todo-go/models"
)

func GetTasks(c *gin.Context) {
	c.JSON(http.StatusOK, models.GetTasks())
}
