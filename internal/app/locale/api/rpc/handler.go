package rpc

import (
	"airbnb-user-be/internal/app/locale/preset/response"
	context "context"

	"github.com/thoas/go-funk"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (h Handler) GetLocales(ctx context.Context, cmd *GetLocalesCmd) (locales *Locales, err error) {
	res, getUserErr := h.Locale.GetLocales(ctx)
	if getUserErr != nil {
		return nil, getUserErr.Error
	}

	data := funk.Map(res, func(data response.Locale) *Locale {
		return h.mapLocale(data)
	}).([]*Locale)
	locales.Data = data

	return
}

func (h Handler) mapLocale(data response.Locale) (locale *Locale) {
	locale = &Locale{}

	locale.Code = data.Code
	locale.Name = data.Name
	if data.Local != nil {
		locale.Local = *data.Local
	}
	if data.Location != nil {
		locale.Location = *data.Location
	}
	if data.Lcid != nil {
		locale.Lcid = int32(*data.Lcid)
	}
	if data.ISO639_1 != nil {
		locale.Iso639_1 = *data.ISO639_1
	}
	if data.ISO639_2 != nil {
		locale.Iso639_2 = *data.ISO639_2
	}
	locale.CreatedAt = timestamppb.New(data.CreatedAt)
	locale.UpdatedAt = timestamppb.New(data.UpdatedAt)

	return
}
