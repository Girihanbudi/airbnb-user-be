package usecaseimpl

import (
	"airbnb-user-be/internal/app/locale/preset/request"
	"airbnb-user-be/internal/app/locale/preset/response"
	"airbnb-user-be/internal/pkg/stderror"
	"context"
)

func (u Usecase) GetRegion(ctx context.Context, cmd request.GetRegion) (res response.GetRegion, err *stderror.StdError) {

	region, getRegionErr := u.RegionRepo.GetRegion(ctx, cmd.Code)
	if getRegionErr != nil {

	}

	res.Region = region

	return
}
