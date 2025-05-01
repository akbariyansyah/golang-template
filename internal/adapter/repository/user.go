package repository

import (
	"context"
	"database/sql"
	"task_1/internal/domain/user"

	"github.com/jmoiron/sqlx"
)

type userRepo struct {
	db *sqlx.DB
}

func (r *userRepo) Create(ctx context.Context, data *user.User) error {
	query := `INSERT INTO public.user (id, name, email, address, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := r.db.ExecContext(ctx, query, data.ID, data.Name, data.Email, data.Address, data.CreatedAt, data.UpdatedAt)
	return err
}

func (r *userRepo) Update(ctx context.Context, data *user.User) error {
	query := `
		UPDATE public.user 
		SET name = $1, email = $2, address = $3, updated_at = $4 
		WHERE id = $5
	`

	if _, err := r.db.ExecContext(ctx, query,
		data.Name,
		data.Email,
		data.Address,
		data.UpdatedAt,
		data.ID,
	); err != nil {
		return err
	}

	return nil
}

func (r *userRepo) GetByID(ctx context.Context, id string) (*user.User, error) {
	query := `SELECT id, name, email, address, created_at, updated_at FROM public.user WHERE id = $1`
	var result user.User
	err := r.db.GetContext(ctx, &result, query, id)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &result, err
}

func (r *userRepo) List(ctx context.Context) (user.Users, error) {
	query := `SELECT id, name, email, address, created_at, updated_at FROM public.user`
	result := make(user.Users, 0)
	err := r.db.SelectContext(ctx, &result, query)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return result, err
}

func (r *userRepo) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM public.user where id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepo) EmailExists(ctx context.Context, email string) (bool, error) {
	query := `
		SELECT EXISTS (
			SELECT 1 
			FROM public.user 
			WHERE email = $1
		)
	`
	var exists bool
	if err := r.db.QueryRowContext(ctx, query, email).Scan(&exists); err != nil {
		return false, err
	}

	return exists, nil
}

func NewUserRepository(db *sqlx.DB) user.Repository {
	return &userRepo{db: db}
}
