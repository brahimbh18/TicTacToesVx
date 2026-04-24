package match

import "github.com/google/uuid"

type Repository interface {
Create(m Match) (Match, error)
GetByID(id uuid.UUID) (Match, bool)
Update(m Match) bool
}
