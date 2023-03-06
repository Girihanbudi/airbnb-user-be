package locale

type Locale struct {
	Code     string `json:"code"`
	Name     string `json:"name"`
	Local    string `json:"local"`
	Location string `json:"location"`
	LCID     int    `json:"lcid"`
	ISO639_2 string `json:"iso639_2"`
	ISO639_1 string `json:"iso639_1"`
}
