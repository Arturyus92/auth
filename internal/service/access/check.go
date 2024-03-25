package access

import (
	"context"
	"errors"
	"strings"

	"github.com/Arturyus92/auth/internal/utils"
	"google.golang.org/grpc/metadata"
)

var accessibleRoles map[string]int32

// Check - ...
func (s *service) Check(ctx context.Context, address string) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return errors.New("metadata is not provided")
	}

	authHeader, ok := md["authorization"]
	if !ok || len(authHeader) == 0 {
		return errors.New("authorization header is not provided")
	}

	if !strings.HasPrefix(authHeader[0], authPrefix) {
		return errors.New("invalid authorization header format")
	}

	accessToken := strings.TrimPrefix(authHeader[0], authPrefix)

	claims, err := utils.VerifyToken(accessToken, []byte(accessTokenSecretKey))
	if err != nil {
		return errors.New("access token is invalid")
	}

	accessibleMap, err := s.accessibleRoles(ctx)
	if err != nil {
		return errors.New("failed to get accessible roles")
	}

	role, ok := accessibleMap[address]
	if !ok {
		return nil
	}

	if role == claims.Role {
		return nil
	}

	return errors.New("access denied")
}

// Возвращает мапу с адресом эндпоинта и ролью, которая имеет доступ к нему
func (s *service) accessibleRoles(ctx context.Context) (map[string]int32, error) {
	if accessibleRoles == nil {
		accessibleRoles = make(map[string]int32)

		permissions, err := s.permRepository.GetPermission(ctx)
		if err != nil {
			return nil, err
		}

		for _, perm := range permissions {
			accessibleRoles[perm.Permission] = perm.Role
		}
		// Лезем в базу за данными о доступных ролях для каждого эндпоинта
		// Можно кэшировать данные, чтобы не лезть в базу каждый раз

		// Например, для эндпоинта /auth_v1.AuthV1/Get доступна только роль admin
		//accessibleRoles[PATH] = "admin"
	}

	return accessibleRoles, nil
}
