package graph

import (
	"airbnb-user-be/graph/model"
	"airbnb-user-be/internal/app/locale/preset/response"
	"context"
	"fmt"

	"github.com/thoas/go-funk"
)

// CreateLocale is the resolver for the createLocale field.
func (r *mutationResolver) CreateLocale(ctx context.Context, input model.NewLocale) (*model.Locale, error) {
	panic(fmt.Errorf("not implemented: CreateLocale - createLocale"))
}

// RemoveLocale is the resolver for the removeLocale field.
func (r *mutationResolver) RemoveLocale(ctx context.Context, input model.DeleteLocale) (*model.Locale, error) {
	panic(fmt.Errorf("not implemented: RemoveLocale - removeLocale"))
}

// Locales is the resolver for the locales field.
func (r *queryResolver) Locales(ctx context.Context) ([]*model.Locale, error) {
	data, err := r.Resolver.Locale.GetLocales(ctx)
	if err != nil {
		return nil, err
	}

	locales := funk.Map(*data.Locales, func(data response.Locale) *model.Locale {
		var locale model.Locale

		locale.Code = data.Code
		locale.Name = data.Name
		locale.Local = data.Local
		locale.Location = data.Location
		locale.Lcid = data.Lcid
		locale.Iso639_1 = data.ISO639_1
		locale.Iso639_2 = data.ISO639_2

		return &locale
	}).([]*model.Locale)

	return locales, nil
}

// Locale is the resolver for the locale field.
func (r *queryResolver) Locale(ctx context.Context, code string) (*model.Locale, error) {
	data, err := r.Resolver.Locale.GetLocale(ctx, code)
	if err != nil {
		return nil, err
	}

	localeData := data.Locale

	locale := model.Locale{
		Code:     localeData.Code,
		Name:     localeData.Name,
		Local:    localeData.Local,
		Location: localeData.Location,
		Lcid:     localeData.Lcid,
		Iso639_1: localeData.ISO639_1,
		Iso639_2: localeData.ISO639_2,
	}

	return &locale, nil
}
