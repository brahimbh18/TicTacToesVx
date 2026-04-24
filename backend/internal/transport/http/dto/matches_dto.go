package dto

type CreateMatchRequest struct {
Mode         string  `json:"mode"`
BoardSize    int     `json:"boardSize"`
AIDifficulty string  `json:"aiDifficulty"`
InviteeUserID *string `json:"inviteeUserId"`
}

type MoveRequest struct {
CellIndex int `json:"cellIndex"`
}
