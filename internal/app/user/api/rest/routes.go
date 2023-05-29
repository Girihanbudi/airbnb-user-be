package rest

import authmid "airbnb-user-be/internal/app/middleware/auth"

func (h Handler) RegisterApi() {
	user := h.Router.Group("/user")
	{
		user.PUT("/locale", authmid.GinBindUserClaimsIfAny, h.ChangeLocaleSetting)
		user.GET("/currency", authmid.GinBindUserClaimsIfAny, h.ChangeCurrencySetting)
	}
}
