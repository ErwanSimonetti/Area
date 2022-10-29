package routes

import (
	"github.com/gorilla/mux"
	"AREA/pkg/controllers"
)

var AreaRouter = func(router *mux.Router) {
	// router.HandleFunc("/test/", controllers.Test).Methods("POST")
	router.HandleFunc("/register/", controllers.CreateUser).Methods("POST")
	// router.HandleFunc("/register/token", controllers.CreateTokenUser).Methods("POST")
	router.HandleFunc("/login/", controllers.LoginUser).Methods("POST")
	router.HandleFunc("/login/", controllers.GetAllUsers).Methods("GET")
	router.HandleFunc("/login/{userID}", controllers.GetUserById).Methods("GET")
	router.HandleFunc("/login/{userID}", controllers.DeleteUser).Methods("DELETE")


	router.HandleFunc("/discord/auth", controllers.AuthDiscord).Methods("GET")
	router.HandleFunc("/discord/auth/url", controllers.GetDiscordUrl).Methods("GET")

	router.HandleFunc("/weather/auth", controllers.GetWeather).Methods("GET")
}