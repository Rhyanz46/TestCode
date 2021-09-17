package config

import "os"

var Port = ""

func init() {
	Port = os.Getenv("port")
	if Port == "" {
		Port = "8888"
	}
}
