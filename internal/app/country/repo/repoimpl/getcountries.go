package repoimpl

import (
	module "airbnb-user-be/internal/app/country"
	"airbnb-user-be/internal/pkg/pagination"
	"context"
)

func (r Repo) GetCountries(ctx context.Context, paging *pagination.SQLPaging) (countries *[]module.Country, err error) {
	query := r.Gorm.DB.Where("phone_code != ?", 0)

	err = query.
		Scopes(pagination.GormPaginate(&module.Country{}, paging, query)).
		Find(&countries).Error

	return
}
