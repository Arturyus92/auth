package auth

import (
	"context"
	"errors"

	"github.com/Arturyus92/auth/internal/model"
	"github.com/Arturyus92/auth/internal/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetAccessToken-...
func (s *service) GetAccessToken(ctx context.Context, refreshToken string) (string, error) {
	refreshTokenSecretKey, err := s.secretRepository.GetKeyTokens(ctx, refreshTokenName)
	if err != nil {
		return "", errors.New("key receipt error")
	}

	claims, err := utils.VerifyToken(refreshToken, []byte(refreshTokenSecretKey))
	if err != nil {
		return "", status.Errorf(codes.Aborted, "invalid refresh token")
	}

	accessTokenSecretKey, err := s.secretRepository.GetKeyTokens(ctx, accessTokenName)
	if err != nil {
		return "", errors.New("key receipt error")
	}

	accessToken, err := utils.GenerateToken(model.UserClaims{
		Username: claims.Username,
		Role:     claims.Role,
	},
		[]byte(accessTokenSecretKey),
		accessTokenExpiration,
	)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}
