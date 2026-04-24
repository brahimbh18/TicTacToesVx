package ai

import (
	"math"

	"github.com/brahimbh18/tictactoesvx/backend/internal/domain/match"
)

func MediumMove(board []string, size int, aiSymbol string, maxDepth int) int {
	return minimaxBestMove(board, size, aiSymbol, maxDepth)
}

func HardMove(board []string, size int, aiSymbol string) int {
	limit := len(match.LegalMoves(board))
	if size > 4 {
		limit = 5
	}
	return minimaxBestMove(board, size, aiSymbol, limit)
}

func minimaxBestMove(board []string, size int, aiSymbol string, maxDepth int) int {
	moves := match.LegalMoves(board)
	if len(moves) == 0 {
		return -1
	}
	bestScore := math.MinInt
	bestMove := moves[0]
	opp := "O"
	if aiSymbol == "O" {
		opp = "X"
	}
	for _, mv := range moves {
		next := append([]string{}, board...)
		next[mv] = aiSymbol
		s := minimax(next, size, false, aiSymbol, opp, 1, maxDepth)
		if s > bestScore {
			bestScore = s
			bestMove = mv
		}
	}
	return bestMove
}

func minimax(board []string, size int, maximizing bool, me, opp string, depth, maxDepth int) int {
	if w := match.DetectWinner(board, size); w != "" {
		if w == me {
			return 100 - depth
		}
		if w == opp {
			return depth - 100
		}
	}
	if match.IsDraw(board) || depth >= maxDepth {
		return evaluate(board, size, me, opp)
	}
	moves := match.LegalMoves(board)
	if maximizing {
		best := math.MinInt
		for _, mv := range moves {
			next := append([]string{}, board...)
			next[mv] = me
			val := minimax(next, size, false, me, opp, depth+1, maxDepth)
			if val > best {
				best = val
			}
		}
		return best
	}
	best := math.MaxInt
	for _, mv := range moves {
		next := append([]string{}, board...)
		next[mv] = opp
		val := minimax(next, size, true, me, opp, depth+1, maxDepth)
		if val < best {
			best = val
		}
	}
	return best
}

func evaluate(board []string, size int, me, opp string) int {
	score := 0
	for r := 0; r < size; r++ {
		line := make([]string, size)
		for c := 0; c < size; c++ {
			line[c] = board[r*size+c]
		}
		score += scoreLine(line, me, opp)
	}
	for c := 0; c < size; c++ {
		line := make([]string, size)
		for r := 0; r < size; r++ {
			line[r] = board[r*size+c]
		}
		score += scoreLine(line, me, opp)
	}
	d1 := make([]string, size)
	d2 := make([]string, size)
	for i := 0; i < size; i++ {
		d1[i] = board[i*size+i]
		d2[i] = board[i*size+(size-1-i)]
	}
	score += scoreLine(d1, me, opp)
	score += scoreLine(d2, me, opp)
	return score
}
