package match

import "errors"

func ValidateMove(m Match, symbol Symbol, cellIndex int) error {
if m.Status != StatusActive {
return errors.New("match is not active")
}
if m.NextTurn != symbol {
return errors.New("out of turn")
}
if cellIndex < 0 || cellIndex >= len(m.Board) {
return errors.New("cell index out of range")
}
if m.Board[cellIndex] != "" {
return errors.New("illegal move: cell already occupied")
}
return nil
}

func ApplyMove(m Match, symbol Symbol, cellIndex int) Match {
m.Board[cellIndex] = string(symbol)
winner := DetectWinner(m.Board, m.BoardSize)
if winner != "" {
m.Status = StatusFinished
m.Winner = winner
return m
}
if IsDraw(m.Board) {
m.Status = StatusFinished
m.Winner = "draw"
return m
}
if symbol == X {
m.NextTurn = O
} else {
m.NextTurn = X
}
return m
}

func DetectWinner(board []string, size int) string {
lines := make([][]int, 0, size*2+2)
for r := 0; r < size; r++ {
line := make([]int, size)
for c := 0; c < size; c++ {
line[c] = r*size + c
}
lines = append(lines, line)
}
for c := 0; c < size; c++ {
line := make([]int, size)
for r := 0; r < size; r++ {
line[r] = r*size + c
}
lines = append(lines, line)
}
d1 := make([]int, size)
d2 := make([]int, size)
for i := 0; i < size; i++ {
d1[i] = i*size + i
d2[i] = i*size + (size - 1 - i)
}
lines = append(lines, d1, d2)

for _, line := range lines {
first := board[line[0]]
if first == "" {
continue
}
allSame := true
for _, idx := range line[1:] {
if board[idx] != first {
allSame = false
break
}
}
if allSame {
return first
}
}
return ""
}

func IsDraw(board []string) bool {
for _, c := range board {
if c == "" {
return false
}
}
return true
}

func LegalMoves(board []string) []int {
out := make([]int, 0)
for i, c := range board {
if c == "" {
out = append(out, i)
}
}
return out
}
