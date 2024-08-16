package domain

import (
	"context"

	"github.com/harrymuliawan03/go-rest-api/dto"
)

type AuthService interface {
	Login(ctx context.Context, req dto.AuthRequest) (dto.AuthResponse, error)
}
