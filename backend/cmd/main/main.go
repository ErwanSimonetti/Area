package main

import(
	"log"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	// "github.com/gofiber/fiber/middleware/cors"
	"AREA/pkg/routes"
)

func main()  {
	r := mux.NewRouter()
	routes.AreaRouter(r)
	http.Handle("/", r)

	// app.Use(cors.New(cors.Config{
	// 	AllowCredentials: true,
	// }))
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", r))
}
