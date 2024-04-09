package model

// PermissionRepo ...
type PermissionRepo struct {
	Permission string `db:"path"`
	Role       int32  `db:"role"`
}
