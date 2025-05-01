package param

import (
	"task_1/internal/domain/user"
	"time"

	"github.com/google/uuid"
)

type UserCreate struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
}

func (r *UserCreate) ToUser() *user.User {
	return &user.User{
		ID:        uuid.NewString(),
		Name:      r.Name,
		Email:     r.Email,
		Address:   r.Address,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
