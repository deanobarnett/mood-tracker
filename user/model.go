package user

import "time"

// Model is a struct that represents the user data
type Model struct {
	ID                int       `json:"id" db:"id"`
	Email             string    `json:"email" db:"email"`
	EncryptedPassword string    `json:"-" db:"encrypted_password"`
	RememberToken     string    `json:"-" db:"remember_token"`
	ConfirmationToken string    `json:"-" db:"confirmation_token "`
	UpdatedAt         time.Time `json:"updated_at" db:"updated_at"`
	CreatedAt         time.Time `json:"created_at" db:"created_at"`
}
