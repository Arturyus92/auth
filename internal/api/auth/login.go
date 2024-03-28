package auth

import (
	"context"

	"github.com/Arturyus92/auth/internal/converter"
	desc "github.com/Arturyus92/auth/pkg/auth_v1"
)

// Login - ...
func (i *Implementation) Login(ctx context.Context, req *desc.LoginRequest) (*desc.LoginResponse, error) {
	refreshToken, err := i.authService.Login(ctx, converter.ToServiceLogin(req))
	if err != nil {
		return nil, err
	}

	return &desc.LoginResponse{
		RefreshToken: refreshToken,
	}, nil
}
