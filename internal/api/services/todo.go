package services

import (
	"blaze/internal/api/domain"
	"blaze/pkg/models"
	"blaze/pkg/storage"
	"context"
	"errors"

	"github.com/google/uuid"
)

type TodoService struct {
	todoDb *storage.TodoStorage
}

func NewTodoService(todoDb *storage.TodoStorage) *TodoService {
	return &TodoService{
		todoDb: todoDb,
	}
}

func (s *TodoService) GetTodos(ctx context.Context) ([]models.Todo, error) {
	return s.todoDb.GetTodos(ctx)
}

func (s *TodoService) GetTodoById(ctx context.Context, id string) (models.Todo, error) {
	todo, err := s.todoDb.GetTodoById(ctx, id)
	if err != nil {
		if errors.Is(err, storage.ErrTodoNotFound) {
			return models.Todo{}, domain.ErrTodoNotFound
		}
		return models.Todo{}, err
	}

	return todo, nil
}

func (s *TodoService) CreateTodo(ctx context.Context, newTodo models.NewTodo) (models.Todo, error) {
	todo := models.Todo{
		Id:        uuid.New().String(),
		Title:     newTodo.Title,
		Completed: false,
	}

	err := s.todoDb.CreateTodo(ctx, todo)
	if err != nil {
		return models.Todo{}, err
	}

	return todo, nil
}
