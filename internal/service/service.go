package service

import (
	"context"

	"github.com/Arturyus92/auth/internal/model"
)

// UserService - ...
type UserService interface {
	Create(ctx context.Context, user *model.UserToCreate) (int64, error)
	Get(ctx context.Context, id int64) (*model.User, error)
	Update(ctx context.Context, user *model.UserToUpdate) error
	Delete(ctx context.Context, id int64) error
}
