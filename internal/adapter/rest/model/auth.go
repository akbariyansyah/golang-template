package model

type SubmitEmailRequest struct {
	Email string `json:"email" validate:"required,email"`
}
