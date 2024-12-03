package usecase

import (
	"context"

	"github.com/stazoloto/todo/internal/domain"
	"github.com/stazoloto/todo/pkg/logger"
)

type TodoUsecase struct {
	repo domain.TodoRepository
}

func NewTodoUsecase(repo domain.TodoRepository) *TodoUsecase {
	return &TodoUsecase{
		repo: repo,
	}
}

func (u *TodoUsecase) GetAllTodos(ctx context.Context) ([]domain.Todo, error) {
	logger := logger.FromContext(ctx)
	logger.Info(ctx, "Fetching all todos")
	return u.repo.GetAll(ctx)
}

func (u *TodoUsecase) GetTodoByID(ctx context.Context, id int64) (domain.Todo, error) {
	logger := logger.FromContext(ctx)
	logger.Info(ctx, "Fetching todo by ID")
	return u.repo.GetByID(ctx, id)
}

func (u *TodoUsecase) CreateTodo(ctx context.Context, todo domain.Todo) (int64, error) {
	logger := logger.FromContext(ctx)
	logger.Info(ctx, "Creating todo")
	return u.repo.Create(ctx, todo)
}

func (u *TodoUsecase) UpdateTodo(ctx context.Context, todo domain.Todo) error {
	logger := logger.FromContext(ctx)
	logger.Info(ctx, "Updating todo")
	return u.repo.Update(ctx, todo)
}

func (u *TodoUsecase) DeleteTodo(ctx context.Context, id int64) error {
	logger := logger.FromContext(ctx)
	logger.Info(ctx, "Deleting todo")
	return u.repo.Delete(ctx, id)
}
