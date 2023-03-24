package rest

func (h Handler) RegisterApi() {
	sessions := h.Router.Group("/sessions")
	{
		sessions.GET("/google", h.ContinueWithGoogle)
		// sessions.POST("/email", r.ContinueWithEmail)
		// sessions.POST("/phone", r.ContinueWithPhoneNumber)

		oauth := sessions.Group("/oauth")
		{
			oauth.GET("/google", h.OauthGoogleCallback)
		}
	}
}
