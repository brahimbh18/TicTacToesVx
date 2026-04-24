package match

import (
	"errors"

	"github.com/google/uuid"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

type CreateInput struct {
	Mode          Mode
	BoardSize     int
	AIDifficulty  string
	OwnerUserID   uuid.UUID
	InviteeUserID *uuid.UUID
}

func (s *Service) Create(input CreateInput) (Match, error) {
	if input.BoardSize < 3 || input.BoardSize > 6 {
		return Match{}, errors.New("boardSize must be between 3 and 6")
	}
	board := make([]string, input.BoardSize*input.BoardSize)
	for i := range board {
		board[i] = ""
	}
	m := Match{
		ID:           uuid.New(),
		Mode:         input.Mode,
		BoardSize:    input.BoardSize,
		Board:        board,
		NextTurn:     X,
		Status:       StatusActive,
		AIDifficulty: input.AIDifficulty,
		Winner:       "",
	}
	switch input.Mode {
	case ModeOnlineFriend:
		if input.InviteeUserID == nil {
			return Match{}, errors.New("inviteeUserId required for online_friend")
		}
		m.Players = []Player{{UserID: input.OwnerUserID, Symbol: X}, {UserID: *input.InviteeUserID, Symbol: O}}
	case ModeOnlineRandom, ModeLocal:
		m.Players = []Player{{UserID: input.OwnerUserID, Symbol: X}}
		m.Status = StatusWaiting
	case ModeAI:
		m.Players = []Player{{UserID: input.OwnerUserID, Symbol: X}}
	default:
		return Match{}, errors.New("invalid mode")
	}
	return s.repo.Create(m)
}

func (s *Service) Get(id uuid.UUID) (Match, bool) {
	return s.repo.GetByID(id)
}

func (s *Service) Save(m Match) bool {
	return s.repo.Update(m)
}
