package service

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/harrymuliawan03/go-rest-api/domain"
	"github.com/harrymuliawan03/go-rest-api/dto"
	"github.com/harrymuliawan03/go-rest-api/internal/config"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	cnf            config.Config
	userRepository domain.UserRepository
}

func NewAuth(cnf *config.Config, userRepository domain.UserRepository) domain.AuthService {
	return AuthService{
		cnf:            *cnf,
		userRepository: userRepository,
	}
}

// Login implements domain.AuthService.
func (a AuthService) Login(ctx context.Context, req dto.AuthRequest) (dto.AuthResponse, error) {
	user, err := a.userRepository.FindByEmail(ctx, req.Email)
	if err != nil {
		return dto.AuthResponse{}, err
	}

	if user.Id == "" {
		return dto.AuthResponse{}, domain.NewNotFoundError("User")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return dto.AuthResponse{}, errors.New("Authentication Failed")
	}

	claim := jwt.MapClaims{
		"id":  user.Id,
		"exp": time.Now().Add(time.Duration(a.cnf.Jwt.Exp) * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenStr, err := token.SignedString([]byte(a.cnf.Jwt.Key))

	if err != nil {
		return dto.AuthResponse{}, errors.New("Authentication Failed")
	}

	return dto.AuthResponse{Token: tokenStr}, nil
}
