package match

import "github.com/google/uuid"

type Status string

type Mode string

type Symbol string

const (
StatusWaiting  Status = "waiting"
StatusActive   Status = "active"
StatusFinished Status = "finished"

ModeOnlineFriend Mode = "online_friend"
ModeOnlineRandom Mode = "online_random"
ModeAI           Mode = "ai"
ModeLocal        Mode = "local"

X Symbol = "X"
O Symbol = "O"
)

type Player struct {
UserID uuid.UUID `json:"userId"`
Symbol Symbol    `json:"symbol"`
}

type Match struct {
ID           uuid.UUID `json:"matchId"`
Mode         Mode      `json:"mode"`
Status       Status    `json:"status"`
BoardSize    int       `json:"boardSize"`
Board        []string  `json:"board"`
Players      []Player  `json:"players"`
NextTurn     Symbol    `json:"nextTurn"`
Winner       string    `json:"winner"`
AIDifficulty string    `json:"aiDifficulty,omitempty"`
}
