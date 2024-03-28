package access

import (
	"github.com/Arturyus92/auth/internal/service"
	desc "github.com/Arturyus92/auth/pkg/access_v1"
)

// Implementation - ...
type Implementation struct {
	desc.UnimplementedAccessV1Server
	accessService service.AccessService
}

// NewImplementation - ...
func NewImplementation(accessService service.AccessService) *Implementation {
	return &Implementation{
		accessService: accessService,
	}
}
