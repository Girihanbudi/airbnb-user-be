package usecase

import (
	"airbnb-user-be/internal/app/locale/preset/request"
	"airbnb-user-be/internal/app/locale/preset/response"
	"airbnb-user-be/internal/pkg/stderror"
	"context"
)

type IRegion interface {
	GetRegion(ctx context.Context, cmd request.GetRegion) (res response.GetRegion, err *stderror.StdError)
	GetRegions(ctx context.Context, cmd request.GetRegions) (res response.GetRegions, err *stderror.StdError)
}
