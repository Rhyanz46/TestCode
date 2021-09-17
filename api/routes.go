package api

import (
	"TestCode/api/user"
)

func (server *Server) InitializeRoutes() {
	userRoute := user.Routes("/user")
	server.Router.HandleFunc(userRoute.Prefix+"/login", userRoute.HandleUserLogin)
}
