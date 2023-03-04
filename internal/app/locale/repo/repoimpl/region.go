package repoimpl

import (
	"airbnb-user-be/internal/app/locale"
	"airbnb-user-be/internal/pkg/gorm"
	"airbnb-user-be/internal/pkg/pagination"
	"context"
)

type Options struct {
	Gorm *gorm.Engine
}

type Repo struct {
	Options
}

func NewRegionRepo(options Options) *Repo {
	return &Repo{options}
}

func (r Repo) GetRegions(ctx context.Context, page *pagination.SQLPaging) (regions *[]locale.Region, paging *pagination.SQLPaging, err error) {
	err = r.Gorm.DB.Limit(page.Limit).Offset(page.GetOffset()).Find(&regions).Error
	return
}

func (r Repo) GetRegion(ctx context.Context, code string) (region *locale.Region, err error) {
	err = r.Gorm.DB.Where("code = ?", code).First(&region).Error
	return
}
