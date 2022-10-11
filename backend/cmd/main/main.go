package main

import(
	"log"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"AREA/pkg/routes"
)

func main()  {
	r := mux.NewRouter()
	routes.AreaRouter(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", r))
}