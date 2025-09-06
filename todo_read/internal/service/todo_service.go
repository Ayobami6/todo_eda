package service

import (
	"context"

	"github.com/Ayobami6/todo_read/internal/model"
)

type TodoService interface {
	// define service methods here
	GetAllTodos(ctx context.Context) ([]model.Todo, error)
}
