package http

import (
	"airbnb-user-be/internal/pkg/env"
	"net/http"
)

func DefaultSameSite() http.SameSite {
	if env.CONFIG.Stage != string(env.StageLocal) {
		return http.SameSiteStrictMode
	} else {
		return http.SameSiteNoneMode
	}
}
