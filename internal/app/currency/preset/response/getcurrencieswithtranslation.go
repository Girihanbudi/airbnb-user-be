package response

type GetCurrencyWithTranslation struct {
	Currencies *[]CurrencyWithTranslation `json:"currencies"`
}

type CurrencyWithTranslation struct {
	Code   string `json:"code"`
	Symbol string `json:"symbol"`
	Name   string `json:"name"`
}
