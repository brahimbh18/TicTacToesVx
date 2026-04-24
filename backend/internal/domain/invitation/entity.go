package invitation

import "github.com/google/uuid"

type Status string

const (
	StatusPending  Status = "pending"
	StatusAccepted Status = "accepted"
	StatusDeclined Status = "declined"
)

type Invitation struct {
	ID            uuid.UUID
	InviterUserID uuid.UUID
	InviteeUserID uuid.UUID
	BoardSize     int
	Status        Status
}
