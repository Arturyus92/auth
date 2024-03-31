package user

import (
	"context"
	"log"
	"time"

	sq "github.com/Masterminds/squirrel"
	"golang.org/x/crypto/bcrypt"

	"github.com/Arturyus92/auth/internal/model"
	"github.com/Arturyus92/auth/internal/repository/user/converter"
	modelRepo "github.com/Arturyus92/auth/internal/repository/user/model"
	"github.com/Arturyus92/platform_common/pkg/db"
)

const (
	tableName = "auth"

	colUserID    = "user_id"
	colName      = "name"
	colPassword  = "password"
	colEmail     = "email"
	colRole      = "role"
	colCreatedAt = "created_at"
	colUpdatedAt = "updated_at"
)

// Repo - ...
type Repo struct {
	db db.Client
}

// NewRepository - ...
func NewRepository(db db.Client) *Repo {
	return &Repo{db: db}
}

// Get - ...
func (r *Repo) Get(ctx context.Context, filter modelRepo.UserFilter) (*model.User, error) {
	// Делаем запрос на получение записи по username из таблицы auth
	builderSelectOne := sq.Select(colUserID, colName, colEmail, colRole, colCreatedAt, colUpdatedAt, colPassword).
		From(tableName).
		PlaceholderFormat(sq.Dollar)

	if filter.ID != nil {
		builderSelectOne = builderSelectOne.Where(sq.Eq{colUserID: filter.ID}).Limit(1)
	}

	if filter.Name != nil {
		builderSelectOne = builderSelectOne.Where(sq.Eq{colName: filter.Name}).Limit(1)
	}

	query, args, err := builderSelectOne.ToSql()
	if err != nil {
		log.Printf("failed to build query: %v", err)
		return nil, err
	}

	q := db.Query{
		Name:     "user_repository.Get",
		QueryRaw: query,
	}

	var getUser modelRepo.User
	err = r.db.DB().ScanOneContext(ctx, &getUser, q, args...)
	if err != nil {
		log.Printf("failed to ScanOneContext: %v", err)
		return nil, err
	}
	return converter.ToUserFromRepo(&getUser), nil
}

/*
// Get - ...
func (r *Repo) Get(ctx context.Context, id int64) (*model.User, error) {
	// Делаем запрос на получение записи из таблицы auth
	builderSelectOne := sq.Select(colUserID, colName, colEmail, colRole, colCreatedAt, colUpdatedAt).
		From(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{colUserID: id}).
		Limit(1)

	query, args, err := builderSelectOne.ToSql()
	if err != nil {
		log.Printf("failed to build query: %v", err)
		return nil, err
	}

	q := db.Query{
		Name:     "user_repository.Get",
		QueryRaw: query,
	}

	var getUser modelRepo.User
	err = r.db.DB().ScanOneContext(ctx, &getUser, q, args...)
	if err != nil {
		log.Printf("failed to ScanOneContext: %v", err)
		return nil, err
	}
	return converter.ToUserFromRepo(&getUser), nil
}
*/
// Create - ...
func (r *Repo) Create(ctx context.Context, user *model.UserToCreate) (int64, error) {
	//Хэш пароля по DefaultCost
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}
	user.Password = string(hashedPassword)

	// Делаем запрос на вставку записи в таблицу auth
	builderInsert := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(colName, colPassword, colEmail, colRole).
		Values(user.Name, user.Password, user.Email, user.Role).
		Suffix("RETURNING " + colUserID)

	query, args, err := builderInsert.ToSql()
	if err != nil {
		log.Printf("failed to build query: %v", err)
		return 0, err
	}

	q := db.Query{
		Name:     "user_repository.Create",
		QueryRaw: query,
	}

	var userID int64
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&userID)
	if err != nil {
		log.Printf("failed to created user: %v", err)
		return 0, err
	}

	return userID, nil
}

// Update - ...
func (r *Repo) Update(ctx context.Context, user *model.UserToUpdate) error {
	// Делаем запрос на обновление записи в таблице auth
	builderUpdate := sq.Update(tableName).PlaceholderFormat(sq.Dollar)
	if len(user.Name.String) > 0 {
		builderUpdate = builderUpdate.Set(colName, user.Name)
	}
	if len(user.Email.String) > 0 {
		builderUpdate = builderUpdate.Set(colEmail, user.Email)
	}
	if user.Role != 0 {
		builderUpdate = builderUpdate.Set(colRole, user.Role)
	}

	builderUpdate = builderUpdate.Set(colUpdatedAt, time.Now()).
		Where(sq.Eq{colUserID: user.ID})

	query, args, err := builderUpdate.ToSql()
	if err != nil {
		log.Printf("failed to build query: %v", err)
		return err
	}

	q := db.Query{
		Name:     "user_repository.Update",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		log.Printf("failed to updated user: %v", err)
		return err
	}

	return nil
}

// Delete - ...
func (r *Repo) Delete(ctx context.Context, id int64) error {
	// Делаем запрос на получение записи из таблицы auth
	builderSelectOne := sq.Select(colUserID, colName).
		From(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{colUserID: id}).
		Limit(1)

	query, args, err := builderSelectOne.ToSql()
	if err != nil {
		log.Printf("failed to build query: %v", err)
		return err
	}

	q := db.Query{
		Name:     "user_repository.Get_Delete",
		QueryRaw: query,
	}

	var getUser modelRepo.User
	err = r.db.DB().ScanOneContext(ctx, &getUser, q, args...)
	if err != nil {
		log.Printf("failed to ScanOneContext: %v", err)
		return err
	}
	//Если user существует, то удаляем
	builderDelete := sq.Delete(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{colUserID: id})

	query, args, err = builderDelete.ToSql()
	if err != nil {
		log.Printf("failed to build query: %v", err)
		return err
	}

	q = db.Query{
		Name:     "user_repository.Delete",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		log.Printf("failed to deleted user: %v", err)
		return err
	}

	return nil
}
