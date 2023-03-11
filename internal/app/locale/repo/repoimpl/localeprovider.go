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

func NewLocaleRepo(options Options) *Repo {
	return &Repo{options}
}
