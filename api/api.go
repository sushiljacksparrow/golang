package api

import (
	"github.com/gin-gonic/gin"

	"github.com/rest-go/api/todos"
)

func ApplyRoutes(r *gin.Engine) {
	todosRoute := r.Group("/api")
	todos.ApplyRoutes(todosRoute)
}
