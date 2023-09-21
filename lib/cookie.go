package lib

import "net/http"

func setcookie(name string, v string) {

	cookie := http.Cookie{
		Name:     "_cookie",
		Value:    session.Uuid,
		HttpOnly: true,
		Path:     "/",
	}
	http.SetCookie(writer, &cookie)

}
