package response

import "time"

type Locale struct {
	Code     string  `json:"code"`
	Name     string  `json:"name"`
	Local    *string `json:"local"`
	Location *string `json:"location"`
	Lcid     *int    `json:"lcid"`
	ISO639_1 *string `json:"iso639_1"`
	ISO639_2 *string `json:"iso639_2"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
