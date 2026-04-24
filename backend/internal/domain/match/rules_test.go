package match

import "testing"

func TestDetectWinner3x3(t *testing.T) {
board := []string{"X", "X", "X", "", "O", "", "", "O", ""}
if got := DetectWinner(board, 3); got != "X" {
t.Fatalf("expected X, got %s", got)
}
}

func TestDetectWinner4x4Diagonal(t *testing.T) {
board := []string{
"O", "", "", "",
"", "O", "", "",
"", "", "O", "",
"", "", "", "O",
}
if got := DetectWinner(board, 4); got != "O" {
t.Fatalf("expected O, got %s", got)
}
}
