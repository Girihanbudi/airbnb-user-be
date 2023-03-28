package gql

import (
	"airbnb-user-be/internal/app/country/preset/request"
	"airbnb-user-be/internal/app/country/preset/response"
	"context"
)

func (h Handler) GetCountries(ctx context.Context, limit, page *int) (*response.GetCountries, error) {
	var req request.GetCountries
	if limit != nil {
		req.Pagination.Limit = *limit
	}

	if page != nil {
		req.Pagination.Page = *page
	}

	res, err := h.Country.GetCountries(ctx, req)
	if err != nil {
		return nil, err.Error
	}

	return &res, nil
}
