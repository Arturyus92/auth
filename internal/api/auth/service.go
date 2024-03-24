package auth

import (
	"github.com/Arturyus92/auth/internal/service"
	desc "github.com/Arturyus92/auth/pkg/auth_v1"
)

// Implementation - ...
type Implementation struct {
	desc.UnimplementedAuthV1Server
	authService service.AuthService
}

// NewImplementation - ...
func NewImplementation(authService service.AuthService) *Implementation {
	return &Implementation{
		authService: authService,
	}
}
