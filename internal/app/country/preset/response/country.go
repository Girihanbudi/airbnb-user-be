package response

type Country struct {
	Iso       string  `json:"iso"`
	Iso3      *string `json:"iso_3"`
	Name      string  `json:"name"`
	NumCode   *int    `json:"numCode"`
	PhoneCode int     `json:"phoneCode"`
}
