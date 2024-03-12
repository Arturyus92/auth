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
	var name, email sql.NullString
	if user.Name != nil {
		name = sql.NullString{
			String: user.Name.GetValue(),
			Valid:  true,
		}
	}
	if user.Email != nil {
		email = sql.NullString{
			String: user.Email.GetValue(),
			Valid:  true,
		}
	}
	return &model.UserToUpdate{
		ID:    user.Id,
		Name:  name,
		Email: email,
		Role:  int32(user.Role),
	}
}
