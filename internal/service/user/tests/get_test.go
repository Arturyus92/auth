package tests

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"

	"github.com/Arturyus92/auth/internal/model"
	"github.com/Arturyus92/auth/internal/repository"
	modelRepo "github.com/Arturyus92/auth/internal/repository/user/model"
	"github.com/Arturyus92/auth/internal/service/user"
	"github.com/Arturyus92/platform_common/pkg/db"
	"github.com/Arturyus92/platform_common/pkg/db/transaction"
)

func (s *TestSuite) TestGet() {
	s.T().Parallel()
	type userRepositoryMockFunc func() repository.UserRepository
	type logRepositoryMockFunc func() repository.LogRepository
	type transactorMockFunc func() db.Transactor

	type args struct {
		ctx context.Context
		req int64
	}

	var (
		ctx = context.Background()

		id        = gofakeit.Int64()
		name      = gofakeit.Animal()
		email     = gofakeit.Animal()
		password  = gofakeit.Animal()
		role      = gofakeit.Int32()
		createdAt = gofakeit.Date()
		updatedAt = gofakeit.Date()

		repoErr = fmt.Errorf("repo error")

		filter = modelRepo.UserFilter{
			ID: &id,
		}

		res = &model.User{
			ID:        id,
			Name:      name,
			Email:     email,
			Password:  password,
			Role:      role,
			CreatedAt: createdAt,
			UpdatedAt: sql.NullTime{
				Time:  updatedAt,
				Valid: true,
			},
		}
	)

	tests := []struct {
		name               string
		args               args
		want               *model.User
		err                error
		userRepositoryMock userRepositoryMockFunc
		logRepositoryMock  logRepositoryMockFunc
		transactorMock     transactorMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: id,
			},
			want: res,
			err:  nil,
			userRepositoryMock: func() repository.UserRepository {
				return s.userRepositoryMock.GetMock.Expect(ctx, filter).Return(res, nil)

			},
			logRepositoryMock: func() repository.LogRepository {
				return s.logRepositoryMock
			},
			transactorMock: func() db.Transactor {
				return s.transactorMock
			},
		},
		{
			name: "service error case",
			args: args{
				ctx: ctx,
				req: id,
			},
			want: res,
			err:  repoErr,
			userRepositoryMock: func() repository.UserRepository {
				return s.userRepositoryMock.GetMock.Expect(ctx, filter).Return(nil, repoErr)
			},
			logRepositoryMock: func() repository.LogRepository {
				return s.logRepositoryMock
			},
			transactorMock: func() db.Transactor {
				return s.transactorMock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		s.T().Run(
			tt.name, func(t *testing.T) {
				t.Parallel()

				userRepoMock := tt.userRepositoryMock()
				logRepoMock := tt.logRepositoryMock()
				txManagerMock := transaction.NewTransactionManager(tt.transactorMock())
				service := user.NewService(userRepoMock, txManagerMock, logRepoMock)

				res, err := service.Get(tt.args.ctx, tt.args.req)
				if err != nil {
					s.Require().Equal(tt.err, err)
					return
				}
				s.Require().Equal(tt.want, res)
			},
		)
	}
}
