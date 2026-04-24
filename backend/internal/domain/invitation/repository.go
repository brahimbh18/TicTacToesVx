package invitation

import "github.com/google/uuid"

type Repository interface {
Create(inviterUserID, inviteeUserID uuid.UUID, boardSize int) (Invitation, error)
ListIncoming(inviteeUserID uuid.UUID) []Invitation
GetByID(id uuid.UUID) (Invitation, bool)
SetStatus(id uuid.UUID, status Status) bool
}
