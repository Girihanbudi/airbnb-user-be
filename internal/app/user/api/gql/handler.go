package gql

import (
	"airbnb-user-be/env/appcontext"
	"airbnb-user-be/internal/app/user/preset/request"
	"airbnb-user-be/internal/app/user/preset/response"
	"context"
)

func (h Handler) Me(ctx context.Context) (*response.Me, error) {
	userId := appcontext.GetUserId(ctx)
	req := request.Me{
		UserId: *userId,
	}
	res, err := h.User.Me(ctx, req)
	if err != nil {
		return nil, err.Error
	}

	return &res, nil
}
