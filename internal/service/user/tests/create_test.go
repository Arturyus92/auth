package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/Arturyus92/auth/internal/model"
	"github.com/Arturyus92/auth/internal/repository"
	userService "github.com/Arturyus92/auth/internal/service/user"
	"github.com/Arturyus92/platform_common/pkg/db"
	dbMocks "github.com/Arturyus92/platform_common/pkg/db/mocks"
	"github.com/Arturyus92/platform_common/pkg/db/pg"
	"github.com/Arturyus92/platform_common/pkg/db/transaction"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
)

func (s *TestSuite) TestCreate() {
	s.T().Parallel()
	type userRepositoryMockFunc func() repository.UserRepository
	type logRepositoryMockFunc func() repository.LogRepository
	type transactorMockFunc func() db.Transactor

	type args struct {
		ctx context.Context
		req *model.UserToCreate
	}

	var (
		txOpts = pgx.TxOptions{IsoLevel: pgx.ReadCommitted}
		txM    dbMocks.TxMock
		ctx    = context.Background()

		id              = gofakeit.Int64()
		userRole        = gofakeit.Int32()
		name            = gofakeit.Name()
		email           = gofakeit.Email()
		password        = gofakeit.Animal()
		passwordConfirm = ""

		repoErr = errors.New("repo error")
		txError = errors.Wrap(repoErr, "failed executing code inside transaction")

		userToCreate = &model.UserToCreate{
			Name:            name,
			Email:           email,
			Password:        password,
			PasswordConfirm: passwordConfirm,
			Role:            userRole,
		}

		logInfo = &model.Log{
			Text: fmt.Sprintf("User created: %d", id),
		}
	)

	tests := []struct {
		name               string
		args               args
		want               int64
		err                error
		userRepositoryMock userRepositoryMockFunc
		logRepositoryMock  logRepositoryMockFunc
		transactorMock     transactorMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: userToCreate,
			},
			want: id,
			err:  nil,
			userRepositoryMock: func() repository.UserRepository {
				return s.userRepositoryMock.CreateMock.Expect(pg.MakeContextTx(ctx, &txM), userToCreate).Return(id, nil)
			},
			logRepositoryMock: func() repository.LogRepository {
				return s.logRepositoryMock.CreateLogMock.Expect(pg.MakeContextTx(ctx, &txM), logInfo).Return(nil)
			},
			transactorMock: func() db.Transactor {
				return s.transactorMock.BeginTxMock.Expect(ctx, txOpts).Return(&txM, nil)
			},
		},
		{
			name: "service error case",
			args: args{
				ctx: ctx,
				req: userToCreate,
			},
			want: 0,
			err:  txError,
			userRepositoryMock: func() repository.UserRepository {
				return s.userRepositoryMock.CreateMock.Expect(pg.MakeContextTx(ctx, &txM), userToCreate).Return(0, repoErr)
			},
			logRepositoryMock: func() repository.LogRepository {
				return s.logRepositoryMock
			},
			transactorMock: func() db.Transactor {
				return s.transactorMock.BeginTxMock.Expect(ctx, txOpts).Return(&txM, nil)
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
				service := userService.NewService(userRepoMock, txManagerMock, logRepoMock)

				res, err := service.Create(tt.args.ctx, tt.args.req)
				if err != nil {
					s.Require().Equal(tt.err.Error(), err.Error())
				}
				s.Require().Equal(tt.want, res)
			},
		)
	}
}
