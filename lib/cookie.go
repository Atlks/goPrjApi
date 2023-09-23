package lib

import "net/http"

func setcookie(name string, v string, w http.ResponseWriter) {

	cookie := http.Cookie{
		Name:  "_cookie",
		Value: "session.Uuid",

		Path: "/",
	}
	http.SetCookie(w, &cookie)

}
