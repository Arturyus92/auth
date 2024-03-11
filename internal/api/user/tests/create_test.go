package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"

	"github.com/stretchr/testify/require"

	"github.com/Arturyus92/auth/internal/api/user"
	"github.com/Arturyus92/auth/internal/model"
	"github.com/Arturyus92/auth/internal/service"
	serviceMocks "github.com/Arturyus92/auth/internal/service/mocks"
	desc "github.com/Arturyus92/auth/pkg/user_v1"
)

func TestCreate(t *testing.T) {
	t.Parallel()
	type userServiceMockFunc func(mc *minimock.Controller) service.UserService

	type args struct {
		ctx context.Context
		req *desc.CreateRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id                    = gofakeit.Int64()
		name                  = gofakeit.Animal()
		email                 = gofakeit.Email()
		password              = gofakeit.Password(true, true, true, true, true, 9)
		passwordConfirm       = gofakeit.Animal()
		role            int32 = 1

		serviceErr = fmt.Errorf("service error")

		req = &desc.CreateRequest{
			User: &desc.UserToCreate{
				Name:            name,
				Email:           email,
				Password:        password,
				PasswordConfirm: passwordConfirm,
				Role:            desc.Role(role),
			},
		}

		userToCreate = &model.UserToCreate{
			Name:            name,
			Email:           email,
			Password:        password,
			PasswordConfirm: passwordConfirm,
			Role:            role,
		}

		res = &desc.CreateResponse{
			Id: id,
		}
	)

	tests := []struct {
		name            string
		args            args
		want            *desc.CreateResponse
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
				mock.CreateMock.Expect(ctx, userToCreate).Return(id, nil)
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
				mock.CreateMock.Expect(ctx, userToCreate).Return(0, serviceErr)
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

			res, err := api.Create(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, res)
		})
	}
}
