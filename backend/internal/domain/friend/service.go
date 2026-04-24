package friend

import "github.com/google/uuid"

type Service struct {
repo Repository
}

func NewService(repo Repository) *Service {
return &Service{repo: repo}
}

func (s *Service) CreateRequest(fromUserID, toUserID uuid.UUID) (FriendRequest, error) {
return s.repo.CreateRequest(fromUserID, toUserID)
}

func (s *Service) ListIncoming(toUserID uuid.UUID) []FriendRequest {
return s.repo.ListIncoming(toUserID)
}

func (s *Service) Accept(requestID uuid.UUID) bool {
req, ok := s.repo.GetRequestByID(requestID)
if !ok || req.Status != RequestPending {
return false
}
if !s.repo.SetRequestStatus(requestID, RequestAccepted) {
return false
}
s.repo.CreateFriendship(req.FromUserID, req.ToUserID)
return true
}

func (s *Service) Decline(requestID uuid.UUID) bool {
req, ok := s.repo.GetRequestByID(requestID)
if !ok || req.Status != RequestPending {
return false
}
return s.repo.SetRequestStatus(requestID, RequestDeclined)
}

func (s *Service) ListFriends(userID uuid.UUID) []uuid.UUID {
return s.repo.ListFriends(userID)
}
