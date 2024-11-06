package services

import (
	"blaze/pkg/models"
	"errors"

	"github.com/google/uuid"
)

type TodoService struct {
	db []models.Todo
}

func NewTodoService(db *[]models.Todo) *TodoService {
	return &TodoService{
		db: *db,
	}
}

// TODO where to put these mfs
var ErrTodoNotFound = errors.New("todo not found")

func (s *TodoService) GetTodos() []models.Todo {
	return s.db
}

func (s *TodoService) GetTodoById(id string) (models.Todo, error) {
	for _, todo := range s.db {
		if todo.Id == id {
			return todo, nil
		}
	}

	return models.Todo{}, ErrTodoNotFound
}

func (s *TodoService) CreateTodo(newTodo models.NewTodo) models.Todo {
	todo := models.Todo{
		Id:        uuid.New().String(),
		Title:     newTodo.Title,
		Completed: false,
	}

	s.db = append(s.db, todo)
	return todo
}
