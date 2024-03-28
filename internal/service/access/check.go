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

	accessTokenSecretKey, err := s.secretRepository.GetKeyTokens(ctx, accessTokenName)
	if err != nil {
		return errors.New("key receipt error")
	}

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

		// Лезем в базу за данными о доступных ролях для каждого эндпоинта
		permissions, err := s.permRepository.GetPermission(ctx)
		if err != nil {
			return nil, err
		}

		for _, perm := range permissions {
			accessibleRoles[perm.Permission] = perm.Role
		}
	}

	return accessibleRoles, nil
}
