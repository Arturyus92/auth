package user

import (
	"context"

	"github.com/Arturyus92/auth/internal/converter"
	"github.com/Arturyus92/auth/internal/logger"
	desc "github.com/Arturyus92/auth/pkg/user_v1"
	"go.uber.org/zap"
)

// Get - ...
func (i *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	logger.Info("Getting user...", zap.Int64("id ", req.GetId()))

	user, err := i.userService.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &desc.GetResponse{
		User: converter.ToUserFromService(user),
	}, nil
}
