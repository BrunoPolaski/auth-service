package routes

import "github.com/gorilla/mux"

func Init() *mux.Router {
	r := mux.NewRouter()

	auth := r.PathPrefix("/auth")
	auth.HandlerFunc("/signin")
}
