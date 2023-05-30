package rest

import authmid "airbnb-user-be/internal/app/middleware/auth"

func (h Handler) RegisterApi() {
	me := h.Router.Group("/me")
	{
		me.PUT("/locale", authmid.GinBindUserClaimsIfAny, h.ChangeLocaleSetting)
		me.PUT("/currency", authmid.GinBindUserClaimsIfAny, h.ChangeCurrencySetting)
	}
}
