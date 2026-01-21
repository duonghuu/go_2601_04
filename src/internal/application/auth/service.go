package auth

import (
	domain "go_2601_04/internal/domain/user"
	"go_2601_04/internal/utils"
	"go_2601_04/pkg/auth"
)

type Service struct {
	userRepo     domain.Repository
	tokenService auth.TokenService
}

func NewAuthService(userRepo domain.Repository, tokenService auth.TokenService) *Service {
	return &Service{
		userRepo:     userRepo,
		tokenService: tokenService,
	}
}

func (s *Service) Login(email, password string) (string, error) {
	user, err := s.userRepo.GetByEmail(email)
	if err != nil {
		return "", err
	}

	hashPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return "", err
	}

	isValidPassword := utils.CheckPasswordHash(user.Password, hashPassword)

	if !isValidPassword {
		return "", err
	}
	accessToken, err := s.tokenService.GenerateAccessToken(*user)
	if err != nil {
		return "", err
	}
	return accessToken, nil
}

func (s *Service) Logout() {
}
