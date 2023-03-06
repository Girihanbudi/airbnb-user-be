package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.25

import (
	"airbnb-user-be/graph/model"
	"airbnb-user-be/internal/app/locale"
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
	data, err := r.Resolver.Locale.GetLocales(ctx, 0, 0)
	if err != nil {
		return nil, err
	}

	locales := funk.Map(data, func(data locale.Locale) model.Locale {
		return model.Locale{
			Code:     data.Code,
			Name:     data.Name,
			Local:    &data.Local,
			Location: &data.Location,
			Lcid:     data.LCID,
			Iso639_1: &data.ISO639_1,
			Iso639_2: &data.ISO639_2,
		}
	}).([]*model.Locale)

	return locales, nil
}

// Locale is the resolver for the locale field.
func (r *queryResolver) Locale(ctx context.Context, code string) (*model.Locale, error) {
	data, err := r.Resolver.Locale.GetLocale(ctx, code)
	if err != nil {
		return nil, err
	}

	locale := model.Locale{
		Code:     data.Code,
		Name:     data.Name,
		Local:    &data.Local,
		Location: &data.Location,
		Lcid:     data.LCID,
		Iso639_1: &data.ISO639_1,
		Iso639_2: &data.ISO639_2,
	}
	return &locale, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
