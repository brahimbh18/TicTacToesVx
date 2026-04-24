package handlers

import (
	"encoding/json"
	"net/http"

	sharederr "github.com/brahimbh18/tictactoesvx/backend/internal/shared/errors"
	"github.com/brahimbh18/tictactoesvx/backend/internal/transport/http/dto"
	"github.com/google/uuid"
)

func (d *Deps) CreateFriendRequest(w http.ResponseWriter, r *http.Request) {
	uid, ok := userIDFromContext(r)
	if !ok {
		sharederr.WriteError(w, http.StatusUnauthorized, "unauthorized", "unauthorized")
		return
	}
	var req dto.CreateFriendRequestBody
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		badRequest(w, "invalid json")
		return
	}
	toID, err := uuid.Parse(req.ToUserID)
	if err != nil {
		badRequest(w, "invalid toUserId")
		return
	}
	fr, err := d.Friends.CreateRequest(uid, toID)
	if err != nil {
		sharederr.WriteError(w, http.StatusBadRequest, "friend_request_error", err.Error())
		return
	}
	sharederr.WriteJSON(w, http.StatusCreated, map[string]any{"requestId": fr.ID.String(), "status": fr.Status})
}

func (d *Deps) IncomingFriendRequests(w http.ResponseWriter, r *http.Request) {
	uid, ok := userIDFromContext(r)
	if !ok {
		sharederr.WriteError(w, http.StatusUnauthorized, "unauthorized", "unauthorized")
		return
	}
	items := d.Friends.ListIncoming(uid)
	resp := make([]map[string]any, 0, len(items))
	for _, item := range items {
		u, _ := d.Users.Me(item.FromUserID)
		resp = append(resp, map[string]any{
			"requestId": item.ID.String(),
			"fromUser":  map[string]any{"id": u.ID.String(), "username": u.Username},
			"status":    item.Status,
		})
	}
	sharederr.WriteJSON(w, http.StatusOK, resp)
}

func (d *Deps) AcceptFriendRequest(w http.ResponseWriter, r *http.Request) {
	requestID, err := uuid.Parse(r.PathValue("requestId"))
	if err != nil {
		badRequest(w, "invalid requestId")
		return
	}
	if !d.Friends.Accept(requestID) {
		sharederr.WriteError(w, http.StatusBadRequest, "request_error", "cannot accept request")
		return
	}
	sharederr.WriteJSON(w, http.StatusOK, map[string]bool{"ok": true})
}

func (d *Deps) DeclineFriendRequest(w http.ResponseWriter, r *http.Request) {
	requestID, err := uuid.Parse(r.PathValue("requestId"))
	if err != nil {
		badRequest(w, "invalid requestId")
		return
	}
	if !d.Friends.Decline(requestID) {
		sharederr.WriteError(w, http.StatusBadRequest, "request_error", "cannot decline request")
		return
	}
	sharederr.WriteJSON(w, http.StatusOK, map[string]bool{"ok": true})
}

func (d *Deps) FriendsList(w http.ResponseWriter, r *http.Request) {
	uid, ok := userIDFromContext(r)
	if !ok {
		sharederr.WriteError(w, http.StatusUnauthorized, "unauthorized", "unauthorized")
		return
	}
	ids := d.Friends.ListFriends(uid)
	resp := make([]map[string]string, 0, len(ids))
	for _, fid := range ids {
		u, ok := d.Users.Me(fid)
		if ok {
			resp = append(resp, map[string]string{"id": u.ID.String(), "username": u.Username})
		}
	}
	sharederr.WriteJSON(w, http.StatusOK, resp)
}
