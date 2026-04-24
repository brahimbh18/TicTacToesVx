package middleware

import (
"context"
"net/http"
"strings"

sharederr "github.com/brahimbh18/tictactoesvx/backend/internal/shared/errors"
"github.com/brahimbh18/tictactoesvx/backend/internal/shared/types"
"github.com/brahimbh18/tictactoesvx/backend/internal/security"
)

func Auth(jwtService *security.JWTService) func(http.Handler) http.Handler {
return func(next http.Handler) http.Handler {
return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
h := strings.TrimSpace(r.Header.Get("Authorization"))
if !strings.HasPrefix(h, "Bearer ") {
sharederr.WriteError(w, http.StatusUnauthorized, "unauthorized", "missing bearer token")
return
}
token := strings.TrimSpace(strings.TrimPrefix(h, "Bearer "))
userID, err := jwtService.ParseToken(token)
if err != nil {
sharederr.WriteError(w, http.StatusUnauthorized, "unauthorized", "invalid token")
return
}
ctx := context.WithValue(r.Context(), types.UserIDContextKey, userID)
next.ServeHTTP(w, r.WithContext(ctx))
})
}
}
