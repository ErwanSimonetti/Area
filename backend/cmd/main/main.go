package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	// "github.com/gofiber/fiber/middleware/cors"
	"AREA/pkg/jobs"
	"AREA/pkg/routes"
	"strconv"
	"strings"
	"github.com/jasonlvhit/gocron"
	// "AREA/pkg/controllers"
)

func action(paramStr string) bool {
	params := strings.Split(paramStr, "\n")
	a, _ := strconv.Atoi(params[0])
	b, _ := strconv.Atoi(params[1])
	fmt.Println("checking", a, " >",  b, "?")
	if a > b {
		return true
	} else {
		return false
	}
}

func reaction(message string) {
	fmt.Println(message)
}

func main()  {
	r := mux.NewRouter()
	routes.AreaRouter(r)
	http.Handle("/", r)

	// storedJobs := []jobs.Job{{
	// 	ActionFunc: action,
	// 	ActionFuncParams: "7\n3",
	// 	ReactionFunc: reaction,
	// 	ReactionFuncParams: "mon caca est plus gros",
		
	// }}
	jobs.CreateNewJob("weather", "discord", "ok", 13)
	jobs.ExecAllJob()
	fmt.Println("wait 4 secs")
	gocron.Every(4).Second().Do(jobs.ExecAllJob)
	<- gocron.Start()
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", r))

}

// func main()  {
// 	controllers.TriggerEachSecondes()
// }

