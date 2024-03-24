package auth

import (
	"time"

	_ "github.com/Arturyus92/auth/internal/config/env"
	"github.com/Arturyus92/auth/internal/repository"
	def "github.com/Arturyus92/auth/internal/service"
)

const (
	refreshTokenSecretKey = "W4/X+LLjehdxptt4YgGFCvMpq5ewptpZZYRHY6A72g0="
	accessTokenSecretKey  = "VqvguGiffXILza1f44TWXowDT4zwf03dtXmqWW4SYyE="

	refreshTokenExpiration = 60 * time.Minute
	accessTokenExpiration  = 2 * time.Minute
)

var _ def.AuthService = (*service)(nil)

type service struct {
	userRepository repository.UserRepository
}

// NewService - ...
func NewService(userRepository repository.UserRepository) *service {
	return &service{
		userRepository: userRepository,
	}
}
