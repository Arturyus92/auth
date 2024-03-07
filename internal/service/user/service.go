package user

import (
	"github.com/Arturyus92/auth/internal/client/db"
	"github.com/Arturyus92/auth/internal/repository"
	def "github.com/Arturyus92/auth/internal/service"
)

var _ def.UserService = (*service)(nil)

type service struct {
	userRepository repository.UserRepository
	txManager      db.TxManager
	logRepository  repository.LogRepository
}

// NewService - ...
func NewService(userRepository repository.UserRepository, txManager db.TxManager, logRepository repository.LogRepository) *service {
	return &service{
		userRepository: userRepository,
		txManager:      txManager,
		logRepository:  logRepository,
	}
}
