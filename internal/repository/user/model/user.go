package model

import (
	"database/sql"
	"time"
)

// User - ...
type User struct {
	ID        int64        `db:"user_id"`
	Name      string       `db:"name"`
	Email     string       `db:"email"`
	Password  string       `db:"password"`
	Role      int32        `db:"role"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}
