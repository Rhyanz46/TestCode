package utils

import (
	"TestCode/config"
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

var Oauth *sessions.Store

func init() {
	store := sessions.NewCookieStore([]byte(config.GoogleSessionSecret))
	store.MaxAge(86400 * 30)
	store.Options.Path = "/"
	store.Options.HttpOnly = true // HttpOnly should always be enabled
	store.Options.Secure = false

	gothic.Store = store
	Oauth = &gothic.Store

	goth.UseProviders(
		google.New(
			config.GoogleClientId,
			config.GoogleSessionSecret,
			config.GoogleOauthCallback,
			"email", "profile",
		),
	)
}
