/** @file main.go
 * @brief main file
 *
 * Where everything begins
 */

// @cond

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
	// "AREA/pkg/models"
)

// @endcond

/** @brief Starts the server and the job's gocron
 * @param r *http.Request
 * @return (string) IP adresse
 */
func main() {
	r := mux.NewRouter()
	routes.AreaRouter(r)
	http.Handle("/", r)
	
	s := gocron.NewScheduler(time.UTC)
	s.Every(5).Seconds().Do(jobs.ExecAllJob)
	s.StartAsync()
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", r))
}
