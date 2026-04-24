package app

import (
"fmt"
"net/http"

"github.com/brahimbh18/tictactoesvx/backend/internal/config"
httptransport "github.com/brahimbh18/tictactoesvx/backend/internal/transport/http"
)

type App struct {
cfg    config.Config
server *http.Server
}

func New() (*App, error) {
cfg := config.Load()
h := httptransport.New(cfg)
return &App{
cfg: cfg,
server: &http.Server{
Addr:    ":" + cfg.Port,
Handler: h,
},
}, nil
}

func (a *App) Run() error {
fmt.Printf("backend listening on :%s\n", a.cfg.Port)
return a.server.ListenAndServe()
}
