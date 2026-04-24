package user

import "github.com/google/uuid"

type Repository interface {
	Create(username, passwordHash string) (User, error)
	GetByUsername(username string) (User, bool)
	GetByID(id uuid.UUID) (User, bool)
	SearchByUsername(query string) []User
}
