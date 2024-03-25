package access

import (
	_ "github.com/Arturyus92/auth/internal/config/env"
	"github.com/Arturyus92/auth/internal/repository"
	def "github.com/Arturyus92/auth/internal/service"
)

const (
	authPrefix           = "Bearer "
	accessTokenSecretKey = "VqvguGiffXILza1f44TWXowDT4zwf03dtXmqWW4SYyE="
)

var _ def.AccessService = (*service)(nil)

type service struct {
	permRepository repository.PermRepository
}

// NewService - ...
func NewService(permRepository repository.PermRepository) *service {
	return &service{
		permRepository: permRepository,
	}
}
