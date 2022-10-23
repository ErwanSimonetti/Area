package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/go-co-op/gocron"
	// "github.com/gofiber/fiber/middleware/cors"

	"AREA/pkg/routes"
	"AREA/pkg/jobs"
)

func main() {
	r := mux.NewRouter()
	routes.AreaRouter(r)
	http.Handle("/", r)
	
	s := gocron.NewScheduler(time.UTC)
	s.Every(5).Seconds().Do(jobs.ExecAllJob)
	s.StartAsync()
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", r))

	// storedJobs := []jobs.Job{{
	// 	ActionFunc: action,
	// 	ActionFuncParams: "7\n3",
	// 	ReactionFunc: reaction,
	// 	ReactionFuncParams: "mon caca est plus gros",

	// }}
	// jobs.CreateNewJob("weather", "discord", "ok", 13)
	// fmt.Println("wait 4 secs")
	// gocron.Every(5).Second().Do(jobs.ExecAllJob)
	// <-gocron.Start()
}

// func main()  {
// 	controllers.TriggerEachSecondes()
// }

