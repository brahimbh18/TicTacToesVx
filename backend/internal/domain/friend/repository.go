package friend

import "github.com/google/uuid"

type Repository interface {
CreateRequest(fromUserID, toUserID uuid.UUID) (FriendRequest, error)
ListIncoming(toUserID uuid.UUID) []FriendRequest
GetRequestByID(id uuid.UUID) (FriendRequest, bool)
SetRequestStatus(id uuid.UUID, status RequestStatus) bool
CreateFriendship(a, b uuid.UUID)
ListFriends(userID uuid.UUID) []uuid.UUID
}
