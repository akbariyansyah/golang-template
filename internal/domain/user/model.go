package user

import "time"

type User struct {
	ID        string    `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	Email     string    `db:"email" json:"email"`
	Address   string    `db:"address" json:"address"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

func NewUser(id, name, email, address string) *User {
	return &User{
		ID:        id,
		Name:      name,
		Email:     email,
		Address:   address,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

type Users []User
