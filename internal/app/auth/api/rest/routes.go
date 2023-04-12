package rest

import authmid "airbnb-user-be/internal/app/middleware/auth"

func (h Handler) RegisterApi() {
	sessions := h.Router.Group("/sessions")
	{
		sessions.GET("/google", authmid.GinValidateNoJwtTokenFound, h.ContinueWithGoogle)
		sessions.GET("/facebook", authmid.GinValidateNoJwtTokenFound, h.ContinueWithFacebook)

		phone := sessions.Group("/phone")
		{
			phone.POST("/initial", authmid.GinValidateNoJwtTokenFound, h.ContinueWithPhone)
			phone.POST("/complete", authmid.GinValidateNoJwtTokenFound, h.CompletePhoneRegistration)
			phone.POST("/generate", authmid.GinValidateNoJwtTokenFound, h.ContinueWithPhone)
		}
		sessions.GET("/phone/initial", authmid.GinValidateNoJwtTokenFound, h.ContinueWithPhone)
		sessions.GET("/phone/complete", authmid.GinValidateNoJwtTokenFound, h.CompletePhoneRegistration)
		sessions.GET("/phone/make", authmid.GinValidateNoJwtTokenFound, h.MakePhoneSession)

		oauth := sessions.Group("/oauth")
		{
			oauth.GET("/google", h.OauthGoogleCallback)
			oauth.GET("/facebook", h.OauthFacebookCallback)
		}
	}
}
