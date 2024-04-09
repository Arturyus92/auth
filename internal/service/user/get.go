package user

import (
	"context"

	"github.com/Arturyus92/auth/internal/model"
	modelRepo "github.com/Arturyus92/auth/internal/repository/user/model"
)

// Get - ...
func (s *service) Get(ctx context.Context, id int64) (*model.User, error) {
	user, err := s.userRepository.Get(ctx, modelRepo.UserFilter{ID: &id})
	if err != nil {
		return nil, err
	}

	return user, nil
}
