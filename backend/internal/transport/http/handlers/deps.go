package handlers

import (
"net/http"

"github.com/brahimbh18/tictactoesvx/backend/internal/domain/ai"
"github.com/brahimbh18/tictactoesvx/backend/internal/domain/friend"
"github.com/brahimbh18/tictactoesvx/backend/internal/domain/invitation"
"github.com/brahimbh18/tictactoesvx/backend/internal/domain/match"
"github.com/brahimbh18/tictactoesvx/backend/internal/domain/user"
sharederr "github.com/brahimbh18/tictactoesvx/backend/internal/shared/errors"
"github.com/brahimbh18/tictactoesvx/backend/internal/shared/types"
"github.com/brahimbh18/tictactoesvx/backend/internal/security"
"github.com/google/uuid"
)

type Deps struct {
Users       *user.Service
Friends     *friend.Service
Invitations *invitation.Service
Matches     *match.Service
AI          *ai.Service
JWT         *security.JWTService
}

func userIDFromContext(r *http.Request) (uuid.UUID, bool) {
v := r.Context().Value(types.UserIDContextKey)
id, ok := v.(uuid.UUID)
return id, ok
}

func badRequest(w http.ResponseWriter, msg string) {
sharederr.WriteError(w, http.StatusBadRequest, "bad_request", msg)
}
