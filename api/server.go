package api

import (
	"TestCode/config"
	"fmt"
	"log"
	"net/http"
)

type Server struct {
	Router *http.ServeMux
}

func (server Server) Run() {
	port := ":" + config.Port
	fmt.Println("Listening to port " + port)
	log.Fatal(http.ListenAndServe(port, server.Router))
}
