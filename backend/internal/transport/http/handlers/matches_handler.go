package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/brahimbh18/tictactoesvx/backend/internal/domain/match"
	sharederr "github.com/brahimbh18/tictactoesvx/backend/internal/shared/errors"
	"github.com/brahimbh18/tictactoesvx/backend/internal/transport/http/dto"
	"github.com/google/uuid"
)

func (d *Deps) CreateMatch(w http.ResponseWriter, r *http.Request) {
	uid, ok := userIDFromContext(r)
	if !ok {
		sharederr.WriteError(w, http.StatusUnauthorized, "unauthorized", "unauthorized")
		return
	}
	var req dto.CreateMatchRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		badRequest(w, "invalid json")
		return
	}
	mode := match.Mode(req.Mode)
	var invitee *uuid.UUID
	if req.InviteeUserID != nil {
		parsed, err := uuid.Parse(*req.InviteeUserID)
		if err != nil {
			badRequest(w, "invalid inviteeUserId")
			return
		}
		invitee = &parsed
	}
	created, err := d.Matches.Create(match.CreateInput{Mode: mode, BoardSize: req.BoardSize, AIDifficulty: req.AIDifficulty, OwnerUserID: uid, InviteeUserID: invitee})
	if err != nil {
		sharederr.WriteError(w, http.StatusBadRequest, "match_error", err.Error())
		return
	}
	sharederr.WriteJSON(w, http.StatusCreated, map[string]any{
		"matchId": created.ID.String(), "status": created.Status, "boardSize": created.BoardSize,
		"players": created.Players, "nextTurn": created.NextTurn,
	})
}

func (d *Deps) GetMatch(w http.ResponseWriter, r *http.Request) {
	matchID, err := uuid.Parse(r.PathValue("matchId"))
	if err != nil {
		badRequest(w, "invalid matchId")
		return
	}
	m, ok := d.Matches.Get(matchID)
	if !ok {
		sharederr.WriteError(w, http.StatusNotFound, "not_found", "match not found")
		return
	}
	winner := any(nil)
	if m.Winner != "" {
		winner = m.Winner
	}
	sharederr.WriteJSON(w, http.StatusOK, map[string]any{
		"matchId": m.ID.String(), "status": m.Status, "boardSize": m.BoardSize,
		"board": m.Board, "nextTurn": m.NextTurn, "winner": winner,
	})
}

func symbolForUser(m match.Match, userID uuid.UUID) match.Symbol {
	for _, p := range m.Players {
		if p.UserID == userID {
			return p.Symbol
		}
	}
	if m.Mode == match.ModeAI || m.Mode == match.ModeLocal || m.Mode == match.ModeOnlineRandom {
		return m.NextTurn
	}
	return ""
}

func (d *Deps) MakeMove(w http.ResponseWriter, r *http.Request) {
	uid, ok := userIDFromContext(r)
	if !ok {
		sharederr.WriteError(w, http.StatusUnauthorized, "unauthorized", "unauthorized")
		return
	}
	matchID, err := uuid.Parse(r.PathValue("matchId"))
	if err != nil {
		badRequest(w, "invalid matchId")
		return
	}
	var req dto.MoveRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		badRequest(w, "invalid json")
		return
	}
	m, ok := d.Matches.Get(matchID)
	if !ok {
		sharederr.WriteError(w, http.StatusNotFound, "not_found", "match not found")
		return
	}
	symbol := symbolForUser(m, uid)
	if symbol == "" {
		sharederr.WriteError(w, http.StatusForbidden, "forbidden", "user not in match")
		return
	}
	if err := match.ValidateMove(m, symbol, req.CellIndex); err != nil {
		sharederr.WriteError(w, http.StatusBadRequest, "illegal_move", err.Error())
		return
	}
	m = match.ApplyMove(m, symbol, req.CellIndex)
	if m.Mode == match.ModeAI && m.Status == match.StatusActive && m.NextTurn == match.O {
		aiMove := d.AI.NextMove(m)
		if aiMove >= 0 {
			if err := match.ValidateMove(m, match.O, aiMove); err == nil {
				m = match.ApplyMove(m, match.O, aiMove)
			}
		}
	}
	d.Matches.Save(m)
	winner := any(nil)
	if m.Winner != "" {
		winner = m.Winner
	}
	sharederr.WriteJSON(w, http.StatusOK, map[string]any{
		"applied":  true,
		"board":    m.Board,
		"nextTurn": m.NextTurn,
		"winner":   winner,
		"status":   m.Status,
	})
}

func (d *Deps) ResignMatch(w http.ResponseWriter, r *http.Request) {
	uid, ok := userIDFromContext(r)
	if !ok {
		sharederr.WriteError(w, http.StatusUnauthorized, "unauthorized", "unauthorized")
		return
	}
	matchID, err := uuid.Parse(r.PathValue("matchId"))
	if err != nil {
		badRequest(w, "invalid matchId")
		return
	}
	m, ok := d.Matches.Get(matchID)
	if !ok {
		sharederr.WriteError(w, http.StatusNotFound, "not_found", "match not found")
		return
	}
	winner := "X"
	if symbolForUser(m, uid) == match.X {
		winner = "O"
	}
	m.Status = match.StatusFinished
	m.Winner = winner
	d.Matches.Save(m)
	sharederr.WriteJSON(w, http.StatusOK, map[string]any{"ok": true, "winner": winner})
}
