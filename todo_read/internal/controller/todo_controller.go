package controller

import (
	"context"
	"time"

	"github.com/Ayobami6/todo_read/internal/service"
	"github.com/gin-gonic/gin"
)

type TodoController struct {
	// inject service here when needed
	todoService service.TodoService
}

func NewTodoController(todoService service.TodoService) *TodoController {
	return &TodoController{
		todoService: todoService,
	}
}

func (tc *TodoController) RegisterRoutes(router *gin.Engine) {
	// define todo-related routes here
	todoRoutes := router.Group("/todos")
	{
		todoRoutes.GET("/", tc.getAllTodos)
	}
}

func (tc *TodoController) getAllTodos(c *gin.Context) {
	// implement the logic to get all todos using the service
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5) // maz 5 seconds timeout
	defer cancel()
	todos, err := tc.todoService.GetAllTodos(ctx)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, todos)
}
