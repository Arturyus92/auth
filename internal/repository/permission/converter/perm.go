package converter

import (
	"github.com/Arturyus92/auth/internal/model"
	modelRepo "github.com/Arturyus92/auth/internal/repository/permission/model"
)

// ToPermFromRepo - ...
func ToPermFromRepo(pathPermissions []*modelRepo.PermissionRepo) []*model.Permission {
	var permissions []*model.Permission
	for _, perm := range pathPermissions {
		permissions = append(permissions, &model.Permission{
			Permission: perm.Permission,
			Role:       perm.Role,
		})
	}
	return permissions
}
