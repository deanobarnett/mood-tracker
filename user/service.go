package user

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	DB *sqlx.DB
}

func NewService(db *sqlx.DB) *Service {
	return &Service{DB: db}
}

const createSQL = `
INSERT INTO users (email, encrypted_password, remember_token, confirmation_token, updated_at, created_at)
	VALUES (:email, :encrypted_password, :remember_token, :confirmation_token, :updated_at, :created_at)
`

func (s *Service) CreateUser(ctx context.Context, email, pass string) (*Model, error) {
	now := time.Now()
	user := &Model{
		Email:         email,
		RememberToken: uuid.New().String(),
		CreatedAt:     now,
		UpdatedAt:     now,
	}

	password, err := bcrypt.GenerateFromPassword([]byte(pass), 10)
	if err != nil {
		return nil, fmt.Errorf("error encrypting password: %v", err)
	}

	user.EncryptedPassword = string(password)

	_, err = s.DB.NamedExec(createSQL, user)
	if err != nil {
		return nil, fmt.Errorf("unable to create entry: %w", err)
	}

	return user, nil
}

func (s *Service) Login(ctx context.Context, email string, password string) (*Model, error) {
	user := &Model{}
	err := s.DB.Get(user, "SELECT * FROM users WHERE email=$1", email)
	if err != nil {
		return nil, fmt.Errorf("could not find user %s: %v", email, err)
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.EncryptedPassword), []byte(token))
	if err != nil {
		return nil, fmt.Errorf("password does not match: %v", err)
	}
	return user, nil
}

func (s *Service) Validate(ctx context.Context, token string) error {
	user := &Model{}
	err := s.DB.Get(user, "SELECT 1 FROM users WHERE remember_token=$1", token)
	if err != nil {
		return nil, fmt.Errorf("could not find user token '%s': %v", token, err)
	}
	return nil
}
