package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/brahimbh18/tictactoesvx/backend/internal/domain/match"
	sharederr "github.com/brahimbh18/tictactoesvx/backend/internal/shared/errors"
	"github.com/brahimbh18/tictactoesvx/backend/internal/transport/http/dto"
	"github.com/google/uuid"
)

func (d *Deps) CreateInvitation(w http.ResponseWriter, r *http.Request) {
	uid, ok := userIDFromContext(r)
	if !ok {
		sharederr.WriteError(w, http.StatusUnauthorized, "unauthorized", "unauthorized")
		return
	}
	var req dto.CreateInvitationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		badRequest(w, "invalid json")
		return
	}
	inviteeID, err := uuid.Parse(req.InviteeUserID)
	if err != nil {
		badRequest(w, "invalid inviteeUserId")
		return
	}
	inv, err := d.Invitations.Create(uid, inviteeID, req.BoardSize)
	if err != nil {
		sharederr.WriteError(w, http.StatusBadRequest, "invitation_error", err.Error())
		return
	}
	sharederr.WriteJSON(w, http.StatusCreated, map[string]any{"invitationId": inv.ID.String(), "status": inv.Status})
}

func (d *Deps) IncomingInvitations(w http.ResponseWriter, r *http.Request) {
	uid, ok := userIDFromContext(r)
	if !ok {
		sharederr.WriteError(w, http.StatusUnauthorized, "unauthorized", "unauthorized")
		return
	}
	items := d.Invitations.Incoming(uid)
	resp := make([]map[string]any, 0, len(items))
	for _, inv := range items {
		u, _ := d.Users.Me(inv.InviterUserID)
		resp = append(resp, map[string]any{
			"invitationId": inv.ID.String(),
			"fromUser":     map[string]string{"id": u.ID.String(), "username": u.Username},
			"boardSize":    inv.BoardSize,
		})
	}
	sharederr.WriteJSON(w, http.StatusOK, resp)
}

func (d *Deps) AcceptInvitation(w http.ResponseWriter, r *http.Request) {
	invID, err := uuid.Parse(r.PathValue("invitationId"))
	if err != nil {
		badRequest(w, "invalid invitationId")
		return
	}
	inv, ok := d.Invitations.Accept(invID)
	if !ok {
		sharederr.WriteError(w, http.StatusBadRequest, "invitation_error", "cannot accept invitation")
		return
	}
	created, err := d.Matches.Create(match.CreateInput{Mode: match.ModeOnlineFriend, BoardSize: inv.BoardSize, OwnerUserID: inv.InviterUserID, InviteeUserID: &inv.InviteeUserID})
	if err != nil {
		sharederr.WriteError(w, http.StatusInternalServerError, "match_error", err.Error())
		return
	}
	sharederr.WriteJSON(w, http.StatusOK, map[string]string{"matchId": created.ID.String()})
}

func (d *Deps) DeclineInvitation(w http.ResponseWriter, r *http.Request) {
	invID, err := uuid.Parse(r.PathValue("invitationId"))
	if err != nil {
		badRequest(w, "invalid invitationId")
		return
	}
	if !d.Invitations.Decline(invID) {
		sharederr.WriteError(w, http.StatusBadRequest, "invitation_error", "cannot decline invitation")
		return
	}
	sharederr.WriteJSON(w, http.StatusOK, map[string]bool{"ok": true})
}
