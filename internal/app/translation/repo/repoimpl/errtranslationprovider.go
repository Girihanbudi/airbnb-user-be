package repoimpl

import (
	"airbnb-user-be/internal/pkg/gorm"
)

type Options struct {
	Gorm *gorm.Engine
}

type Repo struct {
	Options
}

var ErrTranslationRepo Repo

func NewErrTranslationRepo(options Options) *Repo {
	ErrTranslationRepo = Repo{options}
	return &ErrTranslationRepo
}
