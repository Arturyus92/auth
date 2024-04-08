package tests

import (
	"context"
	"database/sql"
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

func (s *TestSuite) TestUpdate() {
	s.T().Parallel()
	type userRepositoryMockFunc func() repository.UserRepository
	type logRepositoryMockFunc func() repository.LogRepository
	type transactorMockFunc func() db.Transactor

	type args struct {
		ctx context.Context
		req *model.UserToUpdate
	}

	var (
		txOpts = pgx.TxOptions{IsoLevel: pgx.ReadCommitted}
		txM    dbMocks.TxMock
		ctx    = context.Background()

		id       = gofakeit.Int64()
		userRole = gofakeit.Int32()
		name     = gofakeit.Name()
		email    = gofakeit.Email()

		repoErr = errors.New("repo error")
		txError = errors.Wrap(repoErr, "failed executing code inside transaction")

		userToUpdate = &model.UserToUpdate{
			ID: id,
			Name: sql.NullString{
				String: name,
				Valid:  true,
			},
			Email: sql.NullString{
				String: email,
				Valid:  true,
			},
			Role: userRole,
		}

		logInfo = &model.Log{
			Text: fmt.Sprintf("User updated: %d", id),
		}
	)

	tests := []struct {
		name               string
		args               args
		want               error
		err                error
		userRepositoryMock userRepositoryMockFunc
		logRepositoryMock  logRepositoryMockFunc
		transactorMock     transactorMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: userToUpdate,
			},
			want: nil,
			err:  nil,
			userRepositoryMock: func() repository.UserRepository {
				return s.userRepositoryMock.UpdateMock.Expect(pg.MakeContextTx(ctx, &txM), userToUpdate).Return(nil)
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
				req: userToUpdate,
			},
			want: nil,
			err:  txError,
			userRepositoryMock: func() repository.UserRepository {
				return s.userRepositoryMock.UpdateMock.Expect(pg.MakeContextTx(ctx, &txM), userToUpdate).Return(repoErr)
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

				err := service.Update(tt.args.ctx, tt.args.req)
				if err != nil {
					s.Require().Equal(tt.err.Error(), err.Error())
				}
			},
		)
	}
}
