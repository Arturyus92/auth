package perm

import (
	"context"
	"log"

	sq "github.com/Masterminds/squirrel"

	"github.com/Arturyus92/auth/internal/model"
	"github.com/Arturyus92/platform_common/pkg/db"
)

const (
	tableName = "permissions"

	colID   = "id"
	colRole = "role"
	colPath = "path"
)

// Repo - ...
type Repo struct {
	db db.Client
}

// NewRepository - ...
func NewRepository(db db.Client) *Repo {
	return &Repo{db: db}
}

// GetPermission - ...
func (r *Repo) GetPermission(ctx context.Context) ([]model.Permission, error) {

	builderSelectOne := sq.Select(colRole, colPath).
		From(tableName).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builderSelectOne.ToSql()
	if err != nil {
		log.Printf("failed to build query: %v", err)
		return nil, err
	}

	q := db.Query{
		Name:     "perm_repository.GetPermission",
		QueryRaw: query,
	}

	rows, err := r.db.DB().QueryContext(ctx, q, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pathPermissions []model.Permission

	for rows.Next() {
		var pathPermission model.Permission

		err := rows.Scan(&pathPermission.Role, &pathPermission.Permission)
		if err != nil {
			log.Printf("failed to Scan: %v", err)
			return nil, err
		}

		pathPermissions = append(pathPermissions, pathPermission)
	}

	return pathPermissions, nil
}
