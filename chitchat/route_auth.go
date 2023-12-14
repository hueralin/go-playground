package main

import (
	"chitchat/data"
	"log"
	"net/http"
)

func authenticate(w http.ResponseWriter, r *http.Request) {
	user, err := data.GetUserByEmail(r.PostFormValue("email"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if user.Password == data.Encrypt(r.PostFormValue("password")) {
		session, err := user.CreateSession()
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		cookie := http.Cookie{
			Name:     "sid",
			Value:    session.Uuid,
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		http.Redirect(w, r, "/login", http.StatusFound)
	}
}
