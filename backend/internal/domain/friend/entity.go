package friend

import "github.com/google/uuid"

type RequestStatus string

const (
	RequestPending  RequestStatus = "pending"
	RequestAccepted RequestStatus = "accepted"
	RequestDeclined RequestStatus = "declined"
)

type FriendRequest struct {
	ID         uuid.UUID
	FromUserID uuid.UUID
	ToUserID   uuid.UUID
	Status     RequestStatus
}

type Friendship struct {
	UserID       uuid.UUID
	FriendUserID uuid.UUID
}
