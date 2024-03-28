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

// AuthService - ...
type AuthService interface {
	Login(ctx context.Context, login *model.Login) (string, error)
	GetRefreshToken(ctx context.Context, refreshToken string) (string, error)
	GetAccessToken(ctx context.Context, accessToken string) (string, error)
}

// AccessService - ...
type AccessService interface {
	Check(ctx context.Context, address string) error
}
