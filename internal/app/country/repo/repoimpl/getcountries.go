package repoimpl

import (
	module "airbnb-user-be/internal/app/country"
	"airbnb-user-be/internal/pkg/pagination"
	"context"
)

func (r Repo) GetCountries(ctx context.Context, paging *pagination.SQLPaging) (countries *[]module.Country, err error) {

	err = r.Gorm.DB.
		Scopes(pagination.GormPaginate(&module.Country{}, paging, r.Gorm.DB)).
		Find(&countries).Error

	return
}
