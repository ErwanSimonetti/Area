package routes

import (
	"github.com/gorilla/mux"
	"github.com/BlanchoMartin/area/pkg/controllers"
)

var my_router = func(router *mux.Router) {
	router.HandleFunc("/login/", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/login", controllers.GetUser).Methods("GET")
	router.HandleFunc("/login", controllers.GetUser).Methods("GET")
}