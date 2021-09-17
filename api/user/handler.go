package user

import "net/http"

type User struct {
	Prefix string
}

func Routes(prefix string) *User {
	return &User{
		Prefix: prefix,
	}
}

func (user *User) HandleUserLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}
