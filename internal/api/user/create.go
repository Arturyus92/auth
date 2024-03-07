package user

import (
	"context"

	"github.com/Arturyus92/auth/internal/converter"
	desc "github.com/Arturyus92/auth/pkg/user_v1"
)

// Create - ...
func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	id, err := i.userService.Create(ctx, converter.ToUserCreateFromDesc(req.GetUser()))
	if err != nil {
		return nil, err
	}

	return &desc.CreateResponse{
		Id: id,
	}, nil
}
