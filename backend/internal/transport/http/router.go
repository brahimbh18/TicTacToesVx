package httptransport

import (
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/brahimbh18/tictactoesvx/backend/internal/config"
	"github.com/brahimbh18/tictactoesvx/backend/internal/domain/ai"
	"github.com/brahimbh18/tictactoesvx/backend/internal/domain/friend"
	"github.com/brahimbh18/tictactoesvx/backend/internal/domain/invitation"
	"github.com/brahimbh18/tictactoesvx/backend/internal/domain/match"
	"github.com/brahimbh18/tictactoesvx/backend/internal/domain/user"
	"github.com/brahimbh18/tictactoesvx/backend/internal/repository/postgres"
	"github.com/brahimbh18/tictactoesvx/backend/internal/security"
	sharederr "github.com/brahimbh18/tictactoesvx/backend/internal/shared/errors"
	"github.com/brahimbh18/tictactoesvx/backend/internal/transport/http/handlers"
	"github.com/brahimbh18/tictactoesvx/backend/internal/transport/http/middleware"
	"golang.org/x/time/rate"
)

type ipLimiter struct {
	mu       sync.Mutex
	limiters map[string]*rate.Limiter
	rps      rate.Limit
}

func newIPLimiter(rps int) *ipLimiter {
	return &ipLimiter{limiters: map[string]*rate.Limiter{}, rps: rate.Limit(rps)}
}

func (l *ipLimiter) allow(ip string) bool {
	l.mu.Lock()
	defer l.mu.Unlock()
	lim, ok := l.limiters[ip]
	if !ok {
		lim = rate.NewLimiter(l.rps, int(l.rps))
		l.limiters[ip] = lim
	}
	return lim.Allow()
}

func authRateLimit(l *ipLimiter, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		host, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			host = r.RemoteAddr
		}
		if !l.allow(host) {
			sharederr.WriteError(w, http.StatusTooManyRequests, "rate_limited", "too many auth requests")
			return
		}
		next(w, r)
	}
}

func withAuth(authMw func(http.Handler) http.Handler, handler http.HandlerFunc) http.Handler {
	return authMw(handler)
}

func New(cfg config.Config) http.Handler {
	store := postgres.NewStore()
	deps := &handlers.Deps{
		Users:       user.NewService(postgres.NewUserRepo(store)),
		Friends:     friend.NewService(postgres.NewFriendRepo(store)),
		Invitations: invitation.NewService(postgres.NewInvitationRepo(store)),
		Matches:     match.NewService(postgres.NewMatchRepo(store)),
		AI:          ai.NewService(),
		JWT:         security.NewJWTService(cfg.JWTSecret),
	}
	authMw := middleware.Auth(deps.JWT)
	limiter := newIPLimiter(cfg.AuthRateLimitRPS)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		sharederr.WriteJSON(w, http.StatusOK, map[string]any{
			"service": "tictactoesvx-backend",
			"status":  "ok",
			"version": "v1",
			"time":    time.Now().UTC().Format(time.RFC3339),
		})
	})

	mux.HandleFunc("POST /api/v1/auth/register", authRateLimit(limiter, deps.Register))
	mux.HandleFunc("POST /api/v1/auth/login", authRateLimit(limiter, deps.Login))
	mux.Handle("GET /api/v1/users/me", withAuth(authMw, deps.Me))
	mux.Handle("GET /api/v1/users/search", withAuth(authMw, deps.SearchUsers))
	mux.Handle("GET /api/v1/friends", withAuth(authMw, deps.FriendsList))
	mux.Handle("POST /api/v1/friends/requests", withAuth(authMw, deps.CreateFriendRequest))
	mux.Handle("GET /api/v1/friends/requests/incoming", withAuth(authMw, deps.IncomingFriendRequests))
	mux.Handle("POST /api/v1/friends/requests/{requestId}/accept", withAuth(authMw, deps.AcceptFriendRequest))
	mux.Handle("POST /api/v1/friends/requests/{requestId}/decline", withAuth(authMw, deps.DeclineFriendRequest))
	mux.Handle("POST /api/v1/invitations", withAuth(authMw, deps.CreateInvitation))
	mux.Handle("GET /api/v1/invitations/incoming", withAuth(authMw, deps.IncomingInvitations))
	mux.Handle("POST /api/v1/invitations/{invitationId}/accept", withAuth(authMw, deps.AcceptInvitation))
	mux.Handle("POST /api/v1/invitations/{invitationId}/decline", withAuth(authMw, deps.DeclineInvitation))
	mux.Handle("POST /api/v1/matches", withAuth(authMw, deps.CreateMatch))
	mux.Handle("GET /api/v1/matches/{matchId}", withAuth(authMw, deps.GetMatch))
	mux.Handle("POST /api/v1/matches/{matchId}/moves", withAuth(authMw, deps.MakeMove))
	mux.Handle("POST /api/v1/matches/{matchId}/resign", withAuth(authMw, deps.ResignMatch))

	var h http.Handler = mux
	h = middleware.Recover(h)
	h = middleware.Logging(h)
	return h
}
