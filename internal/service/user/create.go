package user

import (
	"context"
	"fmt"

	"github.com/Arturyus92/auth/internal/model"
)

// Create - ...
func (s *serv) Create(ctx context.Context, user *model.UserToCreate) (int64, error) {

	var userID int64

	userInfo := model.UserToCreate{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Role:     user.Role,
	}
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		userID, errTx = s.userRepository.Create(ctx, &userInfo)
		if errTx != nil {
			return errTx
		}

		errTx = s.logRepository.CreateLog(ctx, &model.Log{
			Info: fmt.Sprintf("User created: %d", userID),
		})
		if errTx != nil {
			return errTx
		}

		return nil
	})

	if err != nil {
		return 0, err
	}

	return userID, nil
}
