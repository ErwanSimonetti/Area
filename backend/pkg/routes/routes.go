package routes

import (
	"github.com/gorilla/mux"

	"AREA/pkg/controllers"

)

var AreaRouter = func(router *mux.Router) {
	// router.HandleFunc("/test/", controllers.Test).Methods("POST")
	router.HandleFunc("/register/", controllers.CORS(controllers.CreateUser)).Methods("POST")
	// router.HandleFunc("/register/token", controllers.CreateTokenUser).Methods("POST")
	router.HandleFunc("/login/", controllers.CORS(controllers.LoginUser)).Methods("POST")
	router.HandleFunc("/login/", controllers.GetAllUsers).Methods("GET")
	router.HandleFunc("/logout/", controllers.Logout).Methods("GET")
	router.HandleFunc("/login/{userID}", controllers.GetUserById).Methods("GET")
	router.HandleFunc("/login/{userID}", controllers.DeleteUser).Methods("DELETE")


	router.HandleFunc("/discord/auth", controllers.AuthDiscord).Methods("GET")
	router.HandleFunc("/discord/auth/url", controllers.GetDiscordUrl).Methods("GET")

	router.HandleFunc("/area/create", controllers.CORS(controllers.AddJobToUser)).Methods("POST")
	router.HandleFunc("/area/delete/{ID}", controllers.CORS(controllers.GetDiscordUrl)).Methods("DELETE")
}