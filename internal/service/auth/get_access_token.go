package auth

import (
	"context"

	"github.com/Arturyus92/auth/internal/model"
	"github.com/Arturyus92/auth/internal/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *service) GetAccessToken(ctx context.Context, refreshToken string) (string, error) {
	claims, err := utils.VerifyToken(refreshToken, []byte(refreshTokenSecretKey))
	if err != nil {
		return "", status.Errorf(codes.Aborted, "invalid refresh token")
	}

	// Можем слазать в базу или в кэш за доп данными пользователя

	accessToken, err := utils.GenerateToken(model.UserClaims{
		Username: claims.Username,
		// Это пример, в реальности роль должна браться из базы или кэша
		Role: claims.Role,
	},
		[]byte(accessTokenSecretKey),
		accessTokenExpiration,
	)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}
