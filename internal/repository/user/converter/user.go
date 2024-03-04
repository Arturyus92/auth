package converter

import (
	"github.com/Arturyus92/auth/internal/model"
	modelRepo "github.com/Arturyus92/auth/internal/repository/user/model"
)

// ToUserFromRepo - ...
func ToUserFromRepo(user *modelRepo.User) *model.User {
	return &model.User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
