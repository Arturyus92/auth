package converter

import (
	"github.com/Arturyus92/auth/internal/model"
	modelRepo "github.com/Arturyus92/auth/internal/repository/permission/model"
)

// ToPermFromRepo - ...
func ToPermFromRepo(pathPermissions []*modelRepo.PermissionRepo) []*model.Permission {
	var res []*model.Permission
	for _, perm := range pathPermissions {
		res = append(res, &model.Permission{
			Permission: perm.Permission,
			Role:       perm.Role,
		})
	}
	return res
}
