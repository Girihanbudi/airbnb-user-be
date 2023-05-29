package usecase

import (
	"airbnb-user-be/internal/app/user/preset/request"
	"airbnb-user-be/internal/app/user/preset/response"
	"airbnb-user-be/internal/pkg/stderror"
	"context"
)

type IUser interface {
	Me(ctx context.Context, cmd request.Me) (res response.Me, err *stderror.StdError)
	GetUser(ctx context.Context, cmd request.GetUser) (res response.GetUser, err *stderror.StdError)
	CreateUser(ctx context.Context, cmd request.CreateUser) (res response.CreateUser, err *stderror.StdError)
	UpdateUser(ctx context.Context, cmd request.UpdateUser) (err *stderror.StdError)
	VerifiedUser(ctx context.Context, cmd request.Identifier) (err *stderror.StdError)
	ChangeLocaleSetting(ctx context.Context, cmd request.ChangeLocaleSetting) (err *stderror.StdError)
	ChangeCurrencySetting(ctx context.Context, cmd request.ChangeCurrencySetting) (err *stderror.StdError)
}
