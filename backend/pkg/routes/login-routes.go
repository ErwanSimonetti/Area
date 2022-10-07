package routes

import (
	"github.com/gorilla/mux"
	"AREA/pkg/controllers"
)

var AreaRouter = func(router *mux.Router) {
	router.HandleFunc("/login/", controllers.CreateUser).Methods("POST")
	// router.HandleFunc("/login/", controllers.GetAllUsers).Methods("GET")
	router.HandleFunc("/login/{userID}", controllers.GetUserById).Methods("GET")
	router.HandleFunc("/login/{userID}", controllers.DeleteUser).Methods("DELETE")
}