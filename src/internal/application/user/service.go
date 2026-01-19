package user

import domain "go_2601_04/internal/domain/user"

type Service struct {
	repo domain.Repository
}

func NewUserService(repo domain.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(name, email string) (*domain.User, error) {
	user, err := domain.NewUser(name, email)
	if err != nil {
		return nil, err
	}
	return user, s.repo.Create(user)
}

func (s *Service) Update(id uint, name, email string) (*domain.User, error) {
	user, err := domain.NewUser(name, email)
	if err != nil {
		return nil, err
	}
	user.ID = id
	return user, s.repo.Update(user)
}

func (s *Service) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s *Service) Get(id uint) (*domain.User, error) {
	return s.repo.FindByID(id)
}

func (s *Service) List() ([]domain.User, error) {
	return s.repo.FindAll()
}
