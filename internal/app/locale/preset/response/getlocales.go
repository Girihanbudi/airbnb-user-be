package response

type GetLocales struct {
	Locales *[]Locale `json:"locales"`
}
