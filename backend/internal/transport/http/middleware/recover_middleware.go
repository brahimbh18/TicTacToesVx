package middleware

import (
	"net/http"

	sharederr "github.com/brahimbh18/tictactoesvx/backend/internal/shared/errors"
)

func Recover(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if recover() != nil {
				sharederr.WriteError(w, http.StatusInternalServerError, "internal_error", "internal server error")
			}
		}()
		next.ServeHTTP(w, r)
	})
}
