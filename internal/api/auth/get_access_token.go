package auth

import (
	"context"

	_ "github.com/Arturyus92/auth/internal/converter"
	desc "github.com/Arturyus92/auth/pkg/auth_v1"
)

// GetAccessToken - ...
func (i *Implementation) GetAccessToken(ctx context.Context, req *desc.GetAccessTokenRequest) (*desc.GetAccessTokenResponse, error) {
	accessToken, err := i.authService.GetAccessToken(ctx, req.GetRefreshToken())
	if err != nil {
		return nil, err
	}

	return &desc.GetAccessTokenResponse{
		AccessToken: accessToken,
	}, nil
}
