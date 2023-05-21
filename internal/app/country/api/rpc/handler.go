package rpc

import (
	"airbnb-user-be/internal/app/country/preset/request"
	"airbnb-user-be/internal/app/country/preset/response"
	context "context"
)

// func (h Handler) mustEmbedUnimplementedUserServiceServer() {}

func (h Handler) GetCountryByPhoneCode(ctx context.Context, cmd *GetCountryByPhoneCodeCmd) (country *Country, err error) {
	req := request.GetCountryByPhoneCode{
		PhoneCode: int(cmd.Code),
	}
	res, getCountryErr := h.Country.GetCountryByPhoneCode(ctx, req)
	if getCountryErr != nil {
		return nil, getCountryErr.Error
	}

	country = h.mapCountry(res.Country)

	return
}

func (h Handler) mapCountry(data response.Country) (country *Country) {
	country = &Country{}

	country.Iso = data.Iso
	if data.Iso3 != nil {
		country.Iso3 = *data.Iso3
	}
	country.Name = data.Name
	if data.NumCode != nil {
		country.NumCode = int32(*data.NumCode)
	}
	country.PhoneCode = int32(data.PhoneCode)

	return
}
