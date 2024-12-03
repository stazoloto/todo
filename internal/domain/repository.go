package domain

import "context"

type TodoRepository interface {
	GetAll(ctx context.Context) ([]Todo, error)
	GetByID(ctx context.Context, id int64) (Todo, error)
	Create(ctx context.Context,todo Todo) (int64, error)
	Update(ctx context.Context,todo Todo) error
	Delete(ctx context.Context,id int64) error
}
