package usecaseimpl

import (
	"airbnb-user-be/internal/app/locale/repo"
)

type Options struct {
	RegionRepo repo.IRegion
}

type Usecase struct {
	Options
}

func NewRegionUsecase(options Options) *Usecase {
	return &Usecase{options}
}
