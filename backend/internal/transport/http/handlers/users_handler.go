package handlers

import (
"net/http"

sharederr "github.com/brahimbh18/tictactoesvx/backend/internal/shared/errors"
"github.com/brahimbh18/tictactoesvx/backend/internal/transport/http/dto"
)

func (d *Deps) Me(w http.ResponseWriter, r *http.Request) {
uid, ok := userIDFromContext(r)
if !ok {
sharederr.WriteError(w, http.StatusUnauthorized, "unauthorized", "unauthorized")
return
}
u, ok := d.Users.Me(uid)
if !ok {
sharederr.WriteError(w, http.StatusNotFound, "not_found", "user not found")
return
}
sharederr.WriteJSON(w, http.StatusOK, dto.UserResponse{ID: u.ID.String(), Username: u.Username})
}

func (d *Deps) SearchUsers(w http.ResponseWriter, r *http.Request) {
query := r.URL.Query().Get("username")
users := d.Users.Search(query)
resp := make([]dto.UserResponse, 0, len(users))
for _, u := range users {
resp = append(resp, dto.UserResponse{ID: u.ID.String(), Username: u.Username})
}
sharederr.WriteJSON(w, http.StatusOK, resp)
}
