package access

import (
	"github.com/Arturyus92/auth/internal/repository"
	def "github.com/Arturyus92/auth/internal/service"
)

const (
	authPrefix      = "Bearer "
	accessTokenName = "access"
)

var _ def.AccessService = (*service)(nil)

type service struct {
	permRepository   repository.PermRepository
	secretRepository repository.SecretRepository
}

// NewService - ...
func NewService(permRepository repository.PermRepository, secretRepository repository.SecretRepository) *service {
	return &service{
		permRepository:   permRepository,
		secretRepository: secretRepository,
	}
}
