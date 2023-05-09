package rpc

import (
	"airbnb-user-be/internal/app/user/preset/request"
	"airbnb-user-be/internal/app/user/preset/response"
	context "context"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// func (h Handler) mustEmbedUnimplementedUserServiceServer() {}

func (h Handler) GetUser(ctx context.Context, cmd *GetUserCmd) (user *User, err error) {
	req := request.GetUser{
		Id: &cmd.Id,
	}
	res, getUserErr := h.User.GetUser(ctx, req)
	if getUserErr != nil {
		return nil, getUserErr.Error
	}

	user = h.mapUser(res)

	return
}

func (h Handler) GetUserByEmail(ctx context.Context, cmd *GetUserByEmailCmd) (user *User, err error) {
	req := request.GetUser{
		Email: &cmd.Email,
	}
	res, getUserErr := h.User.GetUser(ctx, req)
	if getUserErr != nil {
		return nil, getUserErr.Error
	}

	user = h.mapUser(res)

	return
}

func (h Handler) GetUserByPhone(ctx context.Context, cmd *GetUserByPhoneCmd) (user *User, err error) {
	countryCode := int(cmd.CountryCode)
	req := request.GetUser{
		CountryCode: &countryCode,
		PhoneNumber: &cmd.PhoneNumber,
	}
	res, getUserErr := h.User.GetUser(ctx, req)
	if getUserErr != nil {
		return nil, getUserErr.Error
	}

	user = h.mapUser(res)

	return
}

func (h Handler) CreateUser(ctx context.Context, user *User) (*CreateUserRes, error) {
	req := request.CreateUser{}
	req.FirstName = user.FirstName
	req.LastName = user.LastName
	req.Email = &user.Email
	countryCode := int(user.CountryCode)
	req.CountryCode = &countryCode
	req.PhoneNumber = &user.PhoneNumber
	req.Image = &user.Image
	req.Role = user.Role
	dob := user.DateOfBirth.AsTime()
	req.DateOfBirth = &dob
	req.DefaultSetting.Locale = user.DefaultSetting.Locale
	req.DefaultSetting.Currency = user.DefaultSetting.Currency

	data, err := h.User.CreateUser(ctx, req)
	if err != nil {
		return nil, err.Error
	}

	res := CreateUserRes{
		Id: data.Id,
	}

	return &res, nil
}

func (h Handler) UpdateUser(ctx context.Context, user *User) (*Empty, error) {
	req := request.UpdateUser{}
	req.FirstName = user.FirstName
	req.LastName = user.LastName
	req.DateOfBirth = user.DateOfBirth.AsTime()
	req.DefaultSetting.Locale = user.DefaultSetting.Locale
	req.DefaultSetting.Currency = user.DefaultSetting.Currency

	err := h.User.UpdateUser(ctx, req)
	if err != nil {
		return nil, err.Error
	}

	res := Empty{}

	return &res, nil
}

func (h Handler) mapUser(data response.GetUser) (user *User) {
	user = &User{}

	user.Id = data.Id
	user.FirstName = data.FirstName
	user.FullName = data.FullName
	if data.Email != nil {
		user.Email = *data.Email
	}
	if data.CountryCode != nil {
		user.CountryCode = int32(*data.CountryCode)
	}
	if data.PhoneNumber != nil {
		user.PhoneNumber = *data.PhoneNumber
	}
	if data.Image != nil {
		user.Image = *data.Image
	}
	user.Role = data.Role
	if data.DateOfBirth != nil {
		user.DateOfBirth = timestamppb.New(*data.DateOfBirth)
	}
	user.CreatedAt = timestamppb.New(data.CreatedAt)
	user.UpdatedAt = timestamppb.New(data.UpdatedAt)
	if data.VerifiedAt != nil {
		user.VerifiedAt = timestamppb.New(*data.VerifiedAt)
	}
	if data.DeletedAt != nil {
		user.DeletedAt = timestamppb.New(*data.DeletedAt)
	}

	return
}
