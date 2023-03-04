package repo

import (
	"airbnb-user-be/internal/app/locale"
	"airbnb-user-be/internal/pkg/pagination"
	"context"
)

type IRegion interface {
	GetRegions(ctx context.Context, page *pagination.SQLPaging) (regions *[]locale.Region, paging *pagination.SQLPaging, err error)
	GetRegion(ctx context.Context, code string) (region *locale.Region, err error)
}
