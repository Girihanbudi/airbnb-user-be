package repo

import (
	"airbnb-user-be/internal/app/translation"
	"context"
)

type IErrTranslation interface {
	GetErrTranslation(ctx context.Context, code, localeCode string) (translation *translation.ErrTranslation, err error)
}
