package dto

type CreateInvitationRequest struct {
	InviteeUserID string `json:"inviteeUserId"`
	BoardSize     int    `json:"boardSize"`
}
