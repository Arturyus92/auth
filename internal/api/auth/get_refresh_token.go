package auth

import (
	"context"

	_ "github.com/Arturyus92/auth/internal/converter"
	desc "github.com/Arturyus92/auth/pkg/auth_v1"
)

// GetRefreshToken - ...
func (i *Implementation) GetRefreshToken(ctx context.Context, req *desc.GetRefreshTokenRequest) (*desc.GetRefreshTokenResponse, error) {
	refreshToken, err := i.authService.GetRefreshToken(ctx, req.GetOldRefreshToken())
	if err != nil {
		return nil, err
	}

	return &desc.GetRefreshTokenResponse{
		RefreshToken: refreshToken,
	}, nil
}
