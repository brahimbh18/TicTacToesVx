package invitation

import "github.com/google/uuid"

type Service struct {
repo Repository
}

func NewService(repo Repository) *Service {
return &Service{repo: repo}
}

func (s *Service) Create(inviterUserID, inviteeUserID uuid.UUID, boardSize int) (Invitation, error) {
return s.repo.Create(inviterUserID, inviteeUserID, boardSize)
}

func (s *Service) Incoming(inviteeUserID uuid.UUID) []Invitation {
return s.repo.ListIncoming(inviteeUserID)
}

func (s *Service) Accept(id uuid.UUID) (Invitation, bool) {
inv, ok := s.repo.GetByID(id)
if !ok || inv.Status != StatusPending {
return Invitation{}, false
}
if !s.repo.SetStatus(id, StatusAccepted) {
return Invitation{}, false
}
inv.Status = StatusAccepted
return inv, true
}

func (s *Service) Decline(id uuid.UUID) bool {
inv, ok := s.repo.GetByID(id)
if !ok || inv.Status != StatusPending {
return false
}
return s.repo.SetStatus(id, StatusDeclined)
}
