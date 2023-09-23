package lib

import "net/http"

func set_cookie(name string, v string, w http.ResponseWriter) {

	cookie := http.Cookie{
		Name:  "_cookie",
		Value: "session.Uuid",

		Path: "/",
	}
	http.SetCookie(w, &cookie)

}
