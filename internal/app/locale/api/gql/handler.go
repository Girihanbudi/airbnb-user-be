package gql

import (
	"airbnb-user-be/internal/app/locale/preset/request"
	"airbnb-user-be/internal/pkg/pagination"
	"airbnb-user-be/internal/pkg/stderror"

	"github.com/graphql-go/graphql"
)

func (h Handler) GetRegion(params graphql.ResolveParams) (interface{}, error) {
	code, exist := params.Args["code"]
	if !exist {
		return nil, stderror.DEF_AUTH_401.Error
	}

	cmd := request.GetRegion{
		Code: code.(string),
	}

	res, err := h.Region.GetRegion(params.Context, cmd)

	return res, err.Error
}

func (h Handler) GetRegions(params graphql.ResolveParams) (interface{}, error) {
	limit, exist := params.Args["limit"]
	paging := pagination.DefaultSQLPaging
	if exist {
		paging.Limit = limit.(int)
	}
	page, exist := params.Args["page"]
	if exist {
		paging.Page = page.(int)
	}

	cmd := request.GetRegions{
		Pagination: paging,
	}

	res, err := h.Region.GetRegions(params.Context, cmd)

	return res, err.Error
}
