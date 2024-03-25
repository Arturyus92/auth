package auth

import (
	"context"

	"github.com/Arturyus92/auth/internal/model"
	"github.com/Arturyus92/auth/internal/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *service) GetRefreshToken(ctx context.Context, oldRefreshToken string) (string, error) {
	claims, err := utils.VerifyToken(oldRefreshToken, []byte(refreshTokenSecretKey))
	if err != nil {
		return "", status.Errorf(codes.Aborted, "invalid refresh token")
	}

	// Можем слазать в базу или в кэш за доп данными пользователя

	refreshToken, err := utils.GenerateToken(model.UserClaims{
		Username: claims.Username,
		// Это пример, в реальности роль должна браться из базы или кэша
		Role: claims.Role,
	},
		[]byte(refreshTokenSecretKey),
		refreshTokenExpiration,
	)
	if err != nil {
		return "", err
	}

	return refreshToken, nil
}
