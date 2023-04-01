package rest

import authmid "airbnb-user-be/internal/app/middleware/auth"

func (h Handler) RegisterApi() {
	sessions := h.Router.Group("/sessions")
	{
		sessions.GET("/google", authmid.GinValidateNoJwtTokenFound, h.ContinueWithGoogle)
		sessions.GET("/facebook", authmid.GinValidateNoJwtTokenFound, h.ContinueWithFacebook)

		oauth := sessions.Group("/oauth")
		{
			oauth.GET("/google", h.OauthGoogleCallback)
			oauth.GET("/facebook", h.OauthFacebookCallback)
		}
	}
}
