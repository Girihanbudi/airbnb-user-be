package usecase

import (
	"airbnb-user-be/internal/app/user/preset/request"
	"airbnb-user-be/internal/app/user/preset/response"
	"airbnb-user-be/internal/pkg/stderror"
	"context"
)

type IUser interface {
	Me(ctx context.Context, cmd request.Me) (res response.Me, err *stderror.StdError)
}
