package user

import "github.com/google/uuid"

type Service struct {
repo Repository
}

func NewService(repo Repository) *Service {
return &Service{repo: repo}
}

func (s *Service) Register(username, passwordHash string) (User, error) {
return s.repo.Create(username, passwordHash)
}

func (s *Service) LoginLookup(username string) (User, bool) {
return s.repo.GetByUsername(username)
}

func (s *Service) Me(id uuid.UUID) (User, bool) {
return s.repo.GetByID(id)
}

func (s *Service) Search(username string) []User {
return s.repo.SearchByUsername(username)
}
