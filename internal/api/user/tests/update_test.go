package tests

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"

	"github.com/stretchr/testify/require"

	"github.com/Arturyus92/auth/internal/api/user"
	"github.com/Arturyus92/auth/internal/model"
	"github.com/Arturyus92/auth/internal/service"
	serviceMocks "github.com/Arturyus92/auth/internal/service/mocks"
	desc "github.com/Arturyus92/auth/pkg/user_v1"
)

func TestUpdate(t *testing.T) {
	t.Parallel()
	type userServiceMockFunc func(mc *minimock.Controller) service.UserService

	type args struct {
		ctx context.Context
		req *desc.UpdateRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id          = gofakeit.Int64()
		name        = gofakeit.Animal()
		email       = gofakeit.Email()
		role  int32 = 1

		serviceErr = fmt.Errorf("service error")

		req = &desc.UpdateRequest{
			User: &desc.UserToUpdate{
				Id:    id,
				Name:  wrapperspb.String(name),
				Email: wrapperspb.String(email),
				Role:  desc.Role(role),
			},
		}

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
			Role: role,
		}

		res = &empty.Empty{}
	)

	tests := []struct {
		name            string
		args            args
		want            *emptypb.Empty
		err             error
		userServiceMock userServiceMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: res,
			err:  nil,
			userServiceMock: func(mc *minimock.Controller) service.UserService {
				mock := serviceMocks.NewUserServiceMock(mc)
				mock.UpdateMock.Expect(ctx, userToUpdate).Return(nil)
				return mock
			},
		},
		{
			name: "service error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  serviceErr,
			userServiceMock: func(mc *minimock.Controller) service.UserService {
				mock := serviceMocks.NewUserServiceMock(mc)
				mock.UpdateMock.Expect(ctx, userToUpdate).Return(serviceErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			userServiceMock := tt.userServiceMock(mc)
			api := user.NewImplementation(userServiceMock)

			res, err := api.Update(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, res)
		})
	}
}
