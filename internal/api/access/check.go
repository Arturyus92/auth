package access

import (
	"context"

	desc "github.com/Arturyus92/auth/pkg/access_v1"
	"github.com/golang/protobuf/ptypes/empty"
)

// Check - ...
func (i *Implementation) Check(ctx context.Context, req *desc.CheckRequest) (*empty.Empty, error) {
	err := i.accessService.Check(ctx, req.GetEndpointAddress())
	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}
