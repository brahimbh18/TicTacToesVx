package main

import (
"log"

"github.com/brahimbh18/tictactoesvx/backend/internal/app"
)

func main() {
a, err := app.New()
if err != nil {
log.Fatalf("bootstrap error: %v", err)
}
if err := a.Run(); err != nil {
log.Fatalf("server error: %v", err)
}
}
