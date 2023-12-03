package middlewares

import "net/http"

type AuthMw struct {
	Next http.Handler
}

func (auth *AuthMw) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if auth.Next == nil {
		auth.Next = http.DefaultServeMux
	} else {
		if r.Header.Get("authorization") == "" {
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			auth.Next.ServeHTTP(w, r)
		}
	}
}
