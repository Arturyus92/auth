package model

import (
	"database/sql"
	"time"
)

// User - ...
type User struct {
	ID        int64
	Name      string
	Email     string
	Password  string
	Role      int32
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}

// UserToCreate - ...
type UserToCreate struct {
	Name            string
	Email           string
	Password        string
	PasswordConfirm string
	Role            int32
}

// UserToUpdate - ...
type UserToUpdate struct {
	ID    int64
	Name  string
	Email string
	Role  int32
}
