package user

import (
	"context"

	"github.com/Arturyus92/auth/internal/converter"
	desc "github.com/Arturyus92/auth/pkg/user_v1"
	"github.com/golang/protobuf/ptypes/empty"
)

// Update - ...
func (i *Implementation) Update(ctx context.Context, req *desc.UpdateRequest) (*empty.Empty, error) {
	err := i.userService.Update(ctx, converter.ToUserUpdateFromDesc(req.GetUser()))
	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}
