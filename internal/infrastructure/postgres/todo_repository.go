package postgres

import (
	"context"
	"database/sql"

	"github.com/stazoloto/todo/internal/domain"
	"github.com/stazoloto/todo/pkg/logger"
)

type TodoRepository struct {
	db *sql.DB
}

func NewTodoRepository(db *sql.DB) *TodoRepository {
	return &TodoRepository{db: db}
}

func (r *TodoRepository) GetAll(ctx context.Context) ([]domain.Todo, error) {
	logger := logger.FromContext(ctx)
	logger.Info(ctx, "Fetching all todos from database")

	rows, err := r.db.Query("SELECT id, title, completed FROM todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []domain.Todo
	for rows.Next() {
		var todo domain.Todo
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Completed); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

func (r *TodoRepository) GetByID(ctx context.Context, id int64) (domain.Todo, error) {
	logger := logger.FromContext(ctx)
	logger.Info(ctx, "Fetching todo by ID from database")

	var todo domain.Todo
	err := r.db.QueryRow("SELECT id, title, completed FROM todos WHERE id = $1", id).Scan(&todo.ID, &todo.Title, &todo.Completed)
	if err != nil {
		return domain.Todo{}, err
	}
	return todo, nil
}

func (r *TodoRepository) Create(ctx context.Context, todo domain.Todo) (int64, error) {
	logger := logger.FromContext(ctx)
	logger.Info(ctx, "Creating todo in database")

	var id int64
	err := r.db.QueryRow("INSERT INTO todos (title, completed) VALUES ($1, $2) RETURNING id", todo.Title, todo.Completed).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *TodoRepository) Update(ctx context.Context, todo domain.Todo) error {
	logger := logger.FromContext(ctx)
	logger.Info(ctx, "Updating todo in database")

	_, err := r.db.Exec("UPDATE todos SET title = $1, completed = $2 WHERE id = $3", todo.Title, todo.Completed, todo.ID)
	return err
}

func (r *TodoRepository) Delete(ctx context.Context, id int64) error {
	logger := logger.FromContext(ctx)
	logger.Info(ctx, "Deleting todo from database")

	_, err := r.db.Exec("DELETE FROM todos WHERE id = $1", id)
	return err
}
