package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"

	"github.com/Arturyus92/auth/internal/model"
	"github.com/Arturyus92/auth/internal/repository"
	"github.com/Arturyus92/auth/internal/service/user"
	"github.com/Arturyus92/platform_common/pkg/db"
	dbMocks "github.com/Arturyus92/platform_common/pkg/db/mocks"
	"github.com/Arturyus92/platform_common/pkg/db/pg"
	"github.com/Arturyus92/platform_common/pkg/db/transaction"
)

func (s *TestSuite) TestDelete() {
	s.T().Parallel()
	type userRepositoryMockFunc func() repository.UserRepository
	type logRepositoryMockFunc func() repository.LogRepository
	type transactorMockFunc func() db.Transactor

	type args struct {
		ctx context.Context
		req int64
	}

	var (
		ctx     = context.Background()
		id      = gofakeit.Int64()
		repoErr = fmt.Errorf("repo error")
		txError = errors.Wrap(repoErr, "failed executing code inside transaction")

		txOpts = pgx.TxOptions{IsoLevel: pgx.ReadCommitted}
		txM    dbMocks.TxMock

		logInfo = &model.Log{
			Text: fmt.Sprintf("User deleted: %d", id),
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
				req: id,
			},
			want: nil,
			err:  nil,
			userRepositoryMock: func() repository.UserRepository {
				return s.userRepositoryMock.DeleteMock.Expect(pg.MakeContextTx(ctx, &txM), id).Return(nil)

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
				req: id,
			},
			want: nil,
			err:  txError,
			userRepositoryMock: func() repository.UserRepository {
				return s.userRepositoryMock.DeleteMock.Expect(pg.MakeContextTx(ctx, &txM), id).Return(repoErr)
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
				service := user.NewService(userRepoMock, txManagerMock, logRepoMock)

				err := service.Delete(tt.args.ctx, tt.args.req)
				if err != nil {
					s.Require().Equal(tt.err.Error(), err.Error())
				}
			},
		)
	}
}
