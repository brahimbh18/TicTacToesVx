package ai

import "github.com/brahimbh18/tictactoesvx/backend/internal/domain/match"

type Service struct{}

func NewService() *Service {
return &Service{}
}

func (s *Service) NextMove(m match.Match) int {
legal := match.LegalMoves(m.Board)
if len(legal) == 0 {
return -1
}
switch m.AIDifficulty {
case "hard":
return HardMove(m.Board, m.BoardSize, "O")
case "medium":
return MediumMove(m.Board, m.BoardSize, "O", 4)
default:
return RandomMove(legal)
}
}
