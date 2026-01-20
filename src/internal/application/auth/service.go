package auth

import domain "go_2601_04/internal/domain/user"

type Service struct {
	repo domain.Repository
}

func NewAuthService(repo domain.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Login(email, password string) (*domain.User, error) {
	user, err := s.repo.GetByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Service) Logout() {
}
