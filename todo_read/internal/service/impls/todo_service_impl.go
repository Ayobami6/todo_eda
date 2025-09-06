package impls

import (
	"context"
	"log"

	"github.com/Ayobami6/todo_read/internal/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type TodoServiceImpl struct {
	// add dependencies like DB client here
	db *mongo.Database
}

func NewTodoServiceImpl(db *mongo.Database) *TodoServiceImpl {
	return &TodoServiceImpl{
		db: db,
	}
}

func (s *TodoServiceImpl) GetAllTodos(ctx context.Context) ([]model.Todo, error) {
	// implement the logic to fetch all todos from MongoDB
	collection := s.db.Collection("todos")
	// find all documents in the collection
	cur, err := collection.Find(ctx, map[string]interface{}{})
	if err != nil {
		log.Println("Error fetching todos:", err)
		return nil, err
	}
	var todos []model.Todo
	cur.All(ctx, &todos)
	// return the list of todos
	return todos, nil
}
