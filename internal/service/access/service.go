package access

import (
	_ "github.com/Arturyus92/auth/internal/config/env"
	def "github.com/Arturyus92/auth/internal/service"
)

const (
	authPrefix           = "Bearer "
	accessTokenSecretKey = "VqvguGiffXILza1f44TWXowDT4zwf03dtXmqWW4SYyE="
	PATH                 = "/auth_v1.AuthV1/Get"
)

var _ def.AccessService = (*service)(nil)

type service struct {
}

// NewService - ...
func NewService() *service {
	return &service{}
}
