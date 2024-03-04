package user

import (
	"github.com/Arturyus92/auth/internal/client/db"
	"github.com/Arturyus92/auth/internal/repository"
	"github.com/Arturyus92/auth/internal/service"
)

var _ service.UserService = (*serv)(nil)

type serv struct {
	userRepository repository.UserRepository
	txManager      db.TxManager
	logRepository  repository.LogRepository
}

// NewService - ...
func NewService(userRepository repository.UserRepository, txManager db.TxManager, logRepository repository.LogRepository) *serv {
	return &serv{
		userRepository: userRepository,
		txManager:      txManager,
		logRepository:  logRepository,
	}
}
