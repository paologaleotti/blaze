package storage

import (
	"blaze/pkg/models"
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
)

type TodoStorage struct {
	db *sqlx.DB
}

func NewTodoStorage(db *sql.DB) *TodoStorage {
	client := sqlx.NewDb(db, "sqlite3")
	return &TodoStorage{db: client}
}

func (s *TodoStorage) GetTodos(ctx context.Context) ([]models.Todo, error) {
	query := `SELECT id, title, completed FROM todos`

	todos := []models.Todo{}
	err := s.db.SelectContext(ctx, &todos, query)
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (s *TodoStorage) GetTodoById(ctx context.Context, id string) (models.Todo, error) {
	query := `
		SELECT id, title, completed 
		FROM todos WHERE id = ?
	`

	todo := models.Todo{}
	err := s.db.GetContext(ctx, &todo, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Todo{}, ErrTodoNotFound
		}
		return models.Todo{}, err
	}
	return todo, nil
}

func (s *TodoStorage) CreateTodo(ctx context.Context, newTodo models.Todo) error {
	query := `
		INSERT INTO todos (id, title, completed) 
		VALUES (:id, :title, :completed)
	`

	_, err := s.db.NamedExecContext(ctx, query, &newTodo)
	if err != nil {
		return err
	}
	return nil
}

func (s *TodoStorage) UpdateTodoStatus(ctx context.Context, id string, completed bool) error {
	query := `UPDATE todos SET completed = ? WHERE id = ?`

	_, err := s.db.ExecContext(ctx, query, completed, id)
	return err
}

func (s *TodoStorage) DeleteTodoById(ctx context.Context, id string) error {
	query := `DELETE FROM todos WHERE id = ?`

	_, err := s.db.ExecContext(ctx, query, id)
	return err
}
