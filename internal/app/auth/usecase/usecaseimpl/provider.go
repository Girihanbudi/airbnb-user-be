package usecaseimpl

import (
	countryrepo "airbnb-user-be/internal/app/country/repo"
	localerepo "airbnb-user-be/internal/app/locale/repo"
	userrepo "airbnb-user-be/internal/app/user/repo"
	kafkaproducer "airbnb-user-be/internal/pkg/kafka/producer"
	"airbnb-user-be/internal/pkg/oauth/facebook"
	"airbnb-user-be/internal/pkg/oauth/google"
)

type Options struct {
	GoogleOauth   google.Oauth
	FacebookOauth facebook.Oauth
	UserRepo      userrepo.IUser
	LocaleRepo    localerepo.ILocale
	CountryRepo   countryrepo.ICountry
	EventProducer *kafkaproducer.Producer
}

type Usecase struct {
	Options
}

// Auth Usecase provide a module for authentication
func NewAuthUsecase(options Options) *Usecase {
	return &Usecase{options}
}
