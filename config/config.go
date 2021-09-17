package config

import (
	"fmt"
	"os"
)

var Port = ""
var GoogleClientId = ""
var GoogleSessionSecret = ""
var GoogleOauthCallback = ""

func init() {
	Port = os.Getenv("PORT")
	if Port == "" {
		Port = "8888"
	}
	GoogleSessionSecret = os.Getenv("GOOGLE_SESSION_SECRET")
	if GoogleSessionSecret == "" {
		fmt.Println("you have to set GOOGLE_SESSION_SECRET environment variable")
		os.Exit(1)
	}
}
