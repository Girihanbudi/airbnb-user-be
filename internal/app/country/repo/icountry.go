package repo

import (
	module "airbnb-user-be/internal/app/country"
	"airbnb-user-be/internal/pkg/pagination"
	"context"
)

type ICountry interface {
	GetCountries(ctx context.Context, paging *pagination.SQLPaging) (countries *[]module.Country, err error)
	GetCountryByPhoneCode(ctx context.Context, phoneCode int) (country module.Country, err error)
}
