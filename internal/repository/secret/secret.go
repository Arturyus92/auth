package secret

import (
	"context"
	"log"

	sq "github.com/Masterminds/squirrel"

	"github.com/Arturyus92/platform_common/pkg/db"
)

const (
	tableName = "key_tokens"

	colID    = "id"
	colKey   = "key"
	colValue = "value"
)

// Repo - ...
type Repo struct {
	db db.Client
}

// NewRepository - ...
func NewRepository(db db.Client) *Repo {
	return &Repo{db: db}
}

// GetKeyTokens - ...
func (r *Repo) GetKeyTokens(ctx context.Context, tokenName string) (string, error) {
	builderSelectOne := sq.Select(colValue).
		From(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{colKey: tokenName}).
		Limit(1)

	query, args, err := builderSelectOne.ToSql()
	if err != nil {
		log.Printf("failed to build query: %v", err)
		return "", err
	}

	q := db.Query{
		Name:     "key_tokens_repository.GetKeyTokens",
		QueryRaw: query,
	}

	var value string
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&value)
	if err != nil {
		return "", err
	}

	return value, nil
}
