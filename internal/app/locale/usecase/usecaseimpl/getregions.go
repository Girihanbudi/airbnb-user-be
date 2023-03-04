package usecaseimpl

import (
	"airbnb-user-be/internal/app/locale/preset/request"
	"airbnb-user-be/internal/app/locale/preset/response"
	"airbnb-user-be/internal/pkg/stderror"
	"context"
)

func (u Usecase) GetRegions(ctx context.Context, cmd request.GetRegions) (res response.GetRegions, err *stderror.StdError) {

	regions, paging, getRegionsErr := u.RegionRepo.GetRegions(ctx, &cmd.Pagination)
	if getRegionsErr != nil {

	}

	res.Regions = regions
	res.Paging = paging

	return
}
