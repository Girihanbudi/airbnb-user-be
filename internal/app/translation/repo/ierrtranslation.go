package repo

import (
	module "airbnb-user-be/internal/app/translation"
	"context"
)

type IErrTranslation interface {
	GetErrTranslation(ctx context.Context, code, localeCode string) (translation *module.ErrTranslation, err error)
}
