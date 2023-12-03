package handlers

import "net/http"

func registerHomeHandler() {
	http.HandleFunc("/", home)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Home"))
}
