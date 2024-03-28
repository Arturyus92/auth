package auth

import (
	"context"
	"errors"

	"github.com/Arturyus92/auth/internal/model"
	"github.com/Arturyus92/auth/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

// Login-..
func (s *service) Login(ctx context.Context, login *model.Login) (string, error) {
	// Лезем в базу за данными пользователя
	user, err := s.userRepository.GetLogin(ctx, login.Username)
	if err != nil {
		return "", err
	}

	// Сверяем хэши пароля
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password))
	if err != nil {
		return "", errors.New("invalid password")
	}

	refreshTokenSecretKey, err := s.secretRepository.GetKeyTokens(ctx, refreshTokenName)
	if err != nil {
		return "", errors.New("key receipt error")
	}

	refreshToken, err := utils.GenerateToken(model.UserClaims{
		Username: user.Name,
		Role:     user.Role,
	},
		[]byte(refreshTokenSecretKey),
		refreshTokenExpiration,
	)
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return refreshToken, nil
}
