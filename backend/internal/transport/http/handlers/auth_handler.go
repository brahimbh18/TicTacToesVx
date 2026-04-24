package handlers

import (
"encoding/json"
"net/http"
"strings"

"github.com/brahimbh18/tictactoesvx/backend/internal/security"
sharederr "github.com/brahimbh18/tictactoesvx/backend/internal/shared/errors"
"github.com/brahimbh18/tictactoesvx/backend/internal/transport/http/dto"
)

func (d *Deps) Register(w http.ResponseWriter, r *http.Request) {
var req dto.RegisterRequest
if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
badRequest(w, "invalid json")
return
}
req.Username = strings.TrimSpace(req.Username)
if req.Username == "" || len(req.Password) < 6 {
badRequest(w, "username required and password min 6 chars")
return
}
hash, err := security.HashPassword(req.Password)
if err != nil {
sharederr.WriteError(w, http.StatusInternalServerError, "hash_error", "unable to hash password")
return
}
u, err := d.Users.Register(req.Username, hash)
if err != nil {
sharederr.WriteError(w, http.StatusConflict, "username_taken", err.Error())
return
}
token, err := d.JWT.CreateToken(u.ID)
if err != nil {
sharederr.WriteError(w, http.StatusInternalServerError, "token_error", "unable to create token")
return
}
sharederr.WriteJSON(w, http.StatusCreated, dto.AuthResponse{ID: u.ID.String(), Username: u.Username, Token: token})
}

func (d *Deps) Login(w http.ResponseWriter, r *http.Request) {
var req dto.LoginRequest
if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
badRequest(w, "invalid json")
return
}
u, ok := d.Users.LoginLookup(req.Username)
if !ok || !security.CheckPassword(u.PasswordHash, req.Password) {
sharederr.WriteError(w, http.StatusUnauthorized, "invalid_credentials", "invalid credentials")
return
}
token, err := d.JWT.CreateToken(u.ID)
if err != nil {
sharederr.WriteError(w, http.StatusInternalServerError, "token_error", "unable to create token")
return
}
sharederr.WriteJSON(w, http.StatusOK, dto.AuthResponse{ID: u.ID.String(), Username: u.Username, Token: token})
}
