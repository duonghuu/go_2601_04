package auth

import (
	"go_2601_04/internal/domain/user"
	"go_2601_04/internal/utils"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JWTService struct {
}

type TokenService interface {
	GenerateAccessToken(user user.User) (string, error)
}

type Claims struct {
	UserID string `json:"id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

var jwtSecret = []byte(utils.GetEnv("JWT_SECRET", "Jwt-Learning-Jwt-LearningJwt-Learning-Jwt-Learning"))

const (
	AccessTokenTTL = 15 * time.Minute
)

func NewJWTService() *JWTService {
	return &JWTService{}
}

func (js *JWTService) GenerateAccessToken(user user.User) (string, error) {
	claims := &Claims{
		UserID: strconv.FormatInt(int64(user.ID), 10),
		Email:  user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        uuid.NewString(),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(AccessTokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "eric-le-learner",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
