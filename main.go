package main

import (
	"TestCode/api"
	"net/http"
)

func main() {
	server := api.Server{
		Router: http.NewServeMux(),
	}
	server.InitializeRoutes()
	server.Run()
}
