package main

import(
	"log"
	// "time"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"AREA/pkg/routes"
	// "github.com/gorilla/handlers"
	// "AREA/pkg/controllers"
	// "AREA/pkg/jobs"
)

func main()  {
	r := mux.NewRouter()
	routes.AreaRouter(r)
	http.Handle("/", r)


	// credentials := handlers.AllowCredentials()
	// methods := handlers.AllowedMethods([]string{"GET, POST, PATCH, PUT, DELETE, OPTIONS"})
	// headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	// origins := handlers.AllowedOrigins([]string{"*"})

	// jobs.NewScheduler()

	// gocron.Every(1).Second().Do(task)

	log.Fatal(http.ListenAndServe("0.0.0.0:8080", (r)))
}

// func main()  {
// 	controllers.TriggerEachSecondes()
// }
