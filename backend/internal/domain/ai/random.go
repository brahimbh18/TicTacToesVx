package ai

import (
"math/rand"
"time"
)

func RandomMove(legalMoves []int) int {
if len(legalMoves) == 0 {
return -1
}
rng := rand.New(rand.NewSource(time.Now().UnixNano()))
return legalMoves[rng.Intn(len(legalMoves))]
}
