package user

import (
	"github.com/Arturyus92/auth/internal/service"
	desc "github.com/Arturyus92/auth/pkg/user_v1"
)

// Implementation - ...
type Implementation struct {
	desc.UnimplementedUserV1Server
	userService service.UserService
}

// NewImplementation - ...
func NewImplementation(userService service.UserService) *Implementation {
	return &Implementation{
		userService: userService,
	}
}
