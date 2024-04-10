package auth

import (
	"context"
	"errors"

	"github.com/Arturyus92/auth/internal/model"
	"github.com/Arturyus92/auth/internal/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetRefreshToken-...
func (s *service) GetRefreshToken(ctx context.Context, oldRefreshToken string) (string, error) {
	refreshTokenSecretKey, err := s.secretRepository.GetKeyTokens(ctx, refreshTokenName)
	if err != nil {
		return "", errors.New("key receipt error")
	}

	claims, err := utils.VerifyToken(oldRefreshToken, []byte(refreshTokenSecretKey))
	if err != nil {
		return "", status.Errorf(codes.Aborted, "invalid refresh token")
	}

	refreshToken, err := utils.GenerateToken(model.UserClaims{
		Username: claims.Username,
		Role:     claims.Role,
	},
		[]byte(refreshTokenSecretKey),
		refreshTokenExpiration,
	)
	if err != nil {
		return "", err
	}

	return refreshToken, nil
}
