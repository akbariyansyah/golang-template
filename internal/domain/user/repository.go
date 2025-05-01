package user

import "context"

type Repository interface {
	Create(ctx context.Context, data *User) error
	Update(ctx context.Context, data *User) error
	EmailExists(ctx context.Context, email string) (bool, error)
	GetByID(ctx context.Context, id string) (*User, error)
	List(ctx context.Context) (Users, error)
	Delete(ctx context.Context, id string) error
}
