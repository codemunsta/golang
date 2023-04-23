package middleware

import (
	store "demo/sessions"
	"net/http"
)

func AuthRequired(handler http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		session, _ := store.SessionHouse.Get(request, "session")
		_, ok := session.Values["username"]
		if !ok {
			http.Redirect(writer, request, "/login", 302)
			return
		}
		handler.ServeHTTP(writer, request)
	}
}
