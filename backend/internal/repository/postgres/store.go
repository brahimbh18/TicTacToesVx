package postgres

import (
	"errors"
	"sort"
	"strings"
	"sync"

	"github.com/brahimbh18/tictactoesvx/backend/internal/domain/friend"
	"github.com/brahimbh18/tictactoesvx/backend/internal/domain/invitation"
	"github.com/brahimbh18/tictactoesvx/backend/internal/domain/match"
	"github.com/brahimbh18/tictactoesvx/backend/internal/domain/user"
	"github.com/google/uuid"
)

type Store struct {
	mu sync.RWMutex

	usersByID       map[uuid.UUID]user.User
	usersByUsername map[string]uuid.UUID
	friendRequests  map[uuid.UUID]friend.FriendRequest
	friendships     map[uuid.UUID]map[uuid.UUID]bool
	invitations     map[uuid.UUID]invitation.Invitation
	matches         map[uuid.UUID]match.Match
}

func NewStore() *Store {
	return &Store{
		usersByID:       map[uuid.UUID]user.User{},
		usersByUsername: map[string]uuid.UUID{},
		friendRequests:  map[uuid.UUID]friend.FriendRequest{},
		friendships:     map[uuid.UUID]map[uuid.UUID]bool{},
		invitations:     map[uuid.UUID]invitation.Invitation{},
		matches:         map[uuid.UUID]match.Match{},
	}
}

type UserRepo struct{ store *Store }

func NewUserRepo(store *Store) *UserRepo { return &UserRepo{store: store} }

func (r *UserRepo) Create(username, passwordHash string) (user.User, error) {
	r.store.mu.Lock()
	defer r.store.mu.Unlock()
	u := strings.TrimSpace(strings.ToLower(username))
	if u == "" {
		return user.User{}, errors.New("username is required")
	}
	if _, exists := r.store.usersByUsername[u]; exists {
		return user.User{}, errors.New("username already exists")
	}
	id := uuid.New()
	created := user.User{ID: id, Username: username, PasswordHash: passwordHash}
	r.store.usersByID[id] = created
	r.store.usersByUsername[u] = id
	return created, nil
}

func (r *UserRepo) GetByUsername(username string) (user.User, bool) {
	r.store.mu.RLock()
	defer r.store.mu.RUnlock()
	id, ok := r.store.usersByUsername[strings.TrimSpace(strings.ToLower(username))]
	if !ok {
		return user.User{}, false
	}
	u, ok := r.store.usersByID[id]
	return u, ok
}

func (r *UserRepo) GetByID(id uuid.UUID) (user.User, bool) {
	r.store.mu.RLock()
	defer r.store.mu.RUnlock()
	u, ok := r.store.usersByID[id]
	return u, ok
}

func (r *UserRepo) SearchByUsername(query string) []user.User {
	r.store.mu.RLock()
	defer r.store.mu.RUnlock()
	q := strings.ToLower(strings.TrimSpace(query))
	out := make([]user.User, 0)
	for _, u := range r.store.usersByID {
		if q == "" || strings.Contains(strings.ToLower(u.Username), q) {
			out = append(out, user.User{ID: u.ID, Username: u.Username})
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Username < out[j].Username })
	return out
}

type FriendRepo struct{ store *Store }

func NewFriendRepo(store *Store) *FriendRepo { return &FriendRepo{store: store} }

func (r *FriendRepo) CreateRequest(fromUserID, toUserID uuid.UUID) (friend.FriendRequest, error) {
	r.store.mu.Lock()
	defer r.store.mu.Unlock()
	if fromUserID == toUserID {
		return friend.FriendRequest{}, errors.New("cannot friend yourself")
	}
	for _, req := range r.store.friendRequests {
		if (req.FromUserID == fromUserID && req.ToUserID == toUserID) || (req.FromUserID == toUserID && req.ToUserID == fromUserID) {
			if req.Status == friend.RequestPending {
				return friend.FriendRequest{}, errors.New("request already pending")
			}
		}
	}
	id := uuid.New()
	req := friend.FriendRequest{ID: id, FromUserID: fromUserID, ToUserID: toUserID, Status: friend.RequestPending}
	r.store.friendRequests[id] = req
	return req, nil
}

func (r *FriendRepo) ListIncoming(toUserID uuid.UUID) []friend.FriendRequest {
	r.store.mu.RLock()
	defer r.store.mu.RUnlock()
	out := []friend.FriendRequest{}
	for _, req := range r.store.friendRequests {
		if req.ToUserID == toUserID && req.Status == friend.RequestPending {
			out = append(out, req)
		}
	}
	return out
}

func (r *FriendRepo) GetRequestByID(id uuid.UUID) (friend.FriendRequest, bool) {
	r.store.mu.RLock()
	defer r.store.mu.RUnlock()
	req, ok := r.store.friendRequests[id]
	return req, ok
}

func (r *FriendRepo) SetRequestStatus(id uuid.UUID, status friend.RequestStatus) bool {
	r.store.mu.Lock()
	defer r.store.mu.Unlock()
	req, ok := r.store.friendRequests[id]
	if !ok {
		return false
	}
	req.Status = status
	r.store.friendRequests[id] = req
	return true
}

func (r *FriendRepo) CreateFriendship(a, b uuid.UUID) {
	r.store.mu.Lock()
	defer r.store.mu.Unlock()
	if r.store.friendships[a] == nil {
		r.store.friendships[a] = map[uuid.UUID]bool{}
	}
	if r.store.friendships[b] == nil {
		r.store.friendships[b] = map[uuid.UUID]bool{}
	}
	r.store.friendships[a][b] = true
	r.store.friendships[b][a] = true
}

func (r *FriendRepo) ListFriends(userID uuid.UUID) []uuid.UUID {
	r.store.mu.RLock()
	defer r.store.mu.RUnlock()
	m := r.store.friendships[userID]
	out := make([]uuid.UUID, 0, len(m))
	for id := range m {
		out = append(out, id)
	}
	return out
}

type InvitationRepo struct{ store *Store }

func NewInvitationRepo(store *Store) *InvitationRepo { return &InvitationRepo{store: store} }

func (r *InvitationRepo) Create(inviterUserID, inviteeUserID uuid.UUID, boardSize int) (invitation.Invitation, error) {
	r.store.mu.Lock()
	defer r.store.mu.Unlock()
	id := uuid.New()
	inv := invitation.Invitation{ID: id, InviterUserID: inviterUserID, InviteeUserID: inviteeUserID, BoardSize: boardSize, Status: invitation.StatusPending}
	r.store.invitations[id] = inv
	return inv, nil
}

func (r *InvitationRepo) ListIncoming(inviteeUserID uuid.UUID) []invitation.Invitation {
	r.store.mu.RLock()
	defer r.store.mu.RUnlock()
	out := []invitation.Invitation{}
	for _, inv := range r.store.invitations {
		if inv.InviteeUserID == inviteeUserID && inv.Status == invitation.StatusPending {
			out = append(out, inv)
		}
	}
	return out
}

func (r *InvitationRepo) GetByID(id uuid.UUID) (invitation.Invitation, bool) {
	r.store.mu.RLock()
	defer r.store.mu.RUnlock()
	inv, ok := r.store.invitations[id]
	return inv, ok
}

func (r *InvitationRepo) SetStatus(id uuid.UUID, status invitation.Status) bool {
	r.store.mu.Lock()
	defer r.store.mu.Unlock()
	inv, ok := r.store.invitations[id]
	if !ok {
		return false
	}
	inv.Status = status
	r.store.invitations[id] = inv
	return true
}

type MatchRepo struct{ store *Store }

func NewMatchRepo(store *Store) *MatchRepo { return &MatchRepo{store: store} }

func (r *MatchRepo) Create(m match.Match) (match.Match, error) {
	r.store.mu.Lock()
	defer r.store.mu.Unlock()
	r.store.matches[m.ID] = m
	return m, nil
}

func (r *MatchRepo) GetByID(id uuid.UUID) (match.Match, bool) {
	r.store.mu.RLock()
	defer r.store.mu.RUnlock()
	m, ok := r.store.matches[id]
	return m, ok
}

func (r *MatchRepo) Update(m match.Match) bool {
	r.store.mu.Lock()
	defer r.store.mu.Unlock()
	if _, ok := r.store.matches[m.ID]; !ok {
		return false
	}
	r.store.matches[m.ID] = m
	return true
}
