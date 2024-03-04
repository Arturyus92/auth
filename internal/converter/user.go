package converter

import (
	"database/sql"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/Arturyus92/auth/internal/model"
	desc "github.com/Arturyus92/auth/pkg/user_v1"
)

// ToUserFromService - ...
func ToUserFromService(user *model.User) *desc.User {
	var updatedAt *timestamppb.Timestamp
	if user.UpdatedAt.Valid {
		updatedAt = timestamppb.New(user.UpdatedAt.Time)
	}

	return &desc.User{
		Id:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      desc.Role(user.Role),
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: updatedAt,
	}
}

// ToUserFromDesc - ...
func ToUserFromDesc(user *desc.User) *model.User {
	updatedAt := sql.NullTime{
		Time:  user.UpdatedAt.AsTime(),
		Valid: true,
	}

	return &model.User{
		ID:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		Role:      int32(user.Role),
		CreatedAt: user.CreatedAt.AsTime(),
		UpdatedAt: updatedAt,
	}
}

// ToUserCreateFromDesc - ...
func ToUserCreateFromDesc(user *desc.UserToCreate) *model.UserToCreate {

	return &model.UserToCreate{
		Name:            user.Name,
		Email:           user.Email,
		Password:        user.Password,
		PasswordConfirm: user.PasswordConfirm,
		Role:            int32(user.Role),
	}
}

// ToUserUpdateFromDesc - ...
func ToUserUpdateFromDesc(user *desc.UserToUpdate) *model.UserToUpdate {

	return &model.UserToUpdate{
		ID:    user.Id,
		Name:  user.Name.GetValue(),
		Email: user.Email.GetValue(),
		Role:  int32(user.Role),
	}
}
