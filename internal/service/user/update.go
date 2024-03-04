package user

import (
	"context"
	"fmt"

	"github.com/Arturyus92/auth/internal/model"
)

// Update - ...
func (s *serv) Update(ctx context.Context, user *model.UserToUpdate) error {

	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		errTx := s.userRepository.Update(ctx, user)
		if errTx != nil {
			return errTx
		}

		errTx = s.logRepository.CreateLog(ctx, &model.Log{
			Info: fmt.Sprintf("User updated: %d", user.ID),
		})
		if errTx != nil {
			return errTx
		}
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
