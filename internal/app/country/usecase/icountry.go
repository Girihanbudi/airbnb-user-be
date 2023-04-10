package usecase

import (
	"airbnb-user-be/internal/app/country/preset/request"
	"airbnb-user-be/internal/app/country/preset/response"
	"airbnb-user-be/internal/pkg/stderror"
	"context"
)

type ICountry interface {
	GetCountries(ctx context.Context, cmd request.GetCountries) (res response.GetCountries, err *stderror.StdError)
}
