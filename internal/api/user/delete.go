package user

import (
	"context"

	desc "github.com/Arturyus92/auth/pkg/user_v1"
	"github.com/golang/protobuf/ptypes/empty"
)

// Delete - ...
func (i *Implementation) Delete(ctx context.Context, req *desc.DeleteRequest) (*empty.Empty, error) {
	err := i.userService.Delete(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}
