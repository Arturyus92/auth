package auth

import (
	"time"

	"github.com/Arturyus92/auth/internal/repository"
	def "github.com/Arturyus92/auth/internal/service"
)

const (
	refreshTokenName = "refresh"
	accessTokenName  = "access"

	refreshTokenExpiration = 60 * time.Minute
	accessTokenExpiration  = 2 * time.Minute
)

var _ def.AuthService = (*service)(nil)

type service struct {
	userRepository   repository.UserRepository
	secretRepository repository.SecretRepository
}

// NewService - ...
func NewService(userRepository repository.UserRepository, secretRepository repository.SecretRepository) *service {
	return &service{
		userRepository:   userRepository,
		secretRepository: secretRepository,
	}
}
