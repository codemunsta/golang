package routers

import (
	h "demo/handlers"
	mWare "demo/middleware"
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	fileServer := http.FileServer(http.Dir("./static"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fileServer))
	router.HandleFunc("/", mWare.AuthRequired(h.IndexHandler)).Methods(http.MethodGet)
	router.HandleFunc("/form", mWare.AuthRequired(h.FormHandler)).Methods(http.MethodGet)
	router.HandleFunc("/form", mWare.AuthRequired(h.FormHandler)).Methods(http.MethodPost)
	router.HandleFunc("/login", h.LoginHandler).Methods(http.MethodGet)
	router.HandleFunc("/login", h.LoginHandler).Methods(http.MethodPost)
	router.HandleFunc("/register", h.RegisterHandler).Methods(http.MethodGet)
	router.HandleFunc("/register", h.RegisterHandler).Methods(http.MethodPost)
	// http.HandleFunc("/", handler)
	// router.HandleFunc("/goodbye", goodbyeHandler).Methods(http.MethodGet)
	// router.HandleFunc("/test-session", testHandler).Methods(http.MethodGet)
	return router
}
