package user

import (
	"context"
	"fmt"

	"github.com/Arturyus92/auth/internal/model"
)

// Delete - ...
func (s *service) Delete(ctx context.Context, id int64) error {
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		errTx = s.userRepository.Delete(ctx, id)
		if errTx != nil {
			return errTx
		}

		errTx = s.logRepository.CreateLog(ctx, &model.Log{
			Text: fmt.Sprintf("User deleted: %d", id),
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
