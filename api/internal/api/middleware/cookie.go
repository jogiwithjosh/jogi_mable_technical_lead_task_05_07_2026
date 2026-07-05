package middleware

import (
	"net/http"
	"time"

	"api/internal/config"
)

const (
	AccessTokenCookie = "access_token"
)

func SetAuthCookie(
	w http.ResponseWriter,
	token string,
	cfg config.JWTConfig,
) {
	http.SetCookie(w, &http.Cookie{
		Name:  AccessTokenCookie,
		Value: token,
		Path:  "/",

		HttpOnly: true,

		Secure: false, // localhost

		SameSite: http.SameSiteLaxMode,

		MaxAge: int(cfg.Expiry.Seconds()),

		Expires: time.Now().Add(cfg.Expiry),
	})
}

func ClearAuthCookie(
	w http.ResponseWriter,
) {
	http.SetCookie(w, &http.Cookie{
		Name: AccessTokenCookie,

		Value: "",

		Path: "/",

		HttpOnly: true,

		MaxAge: -1,

		Expires: time.Unix(0, 0),
	})
}
