package repoimpl

import (
	"airbnb-user-be/internal/app/locale"
	"airbnb-user-be/internal/pkg/gorm"
	"airbnb-user-be/internal/pkg/pagination"
	"context"
)

type Options struct {
	Gorm *gorm.Engine
}

type Repo struct {
	Options
}

func NewLocaleRepo(options Options) *Repo {
	return &Repo{options}
}

func (r Repo) GetLocales(ctx context.Context, page *pagination.SQLPaging) (Locales *[]locale.Locale, paging *pagination.SQLPaging, err error) {
	err = r.Gorm.DB.Limit(page.Limit).Offset(page.GetOffset()).Find(&Locales).Error
	return
}

func (r Repo) GetLocale(ctx context.Context, code string) (Locale *locale.Locale, err error) {
	err = r.Gorm.DB.Where("code = ?", code).First(&Locale).Error
	return
}
