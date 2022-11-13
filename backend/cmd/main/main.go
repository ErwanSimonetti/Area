/** @file main.go
 * @brief main file
 *
 * Where everything begins
 */

// @cond

package main

import(
	"log"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	// "github.com/gofiber/fiber/middleware/cors"
	"AREA/pkg/routes"
	// "AREA/pkg/controllers"
)

// @endcond

/** main function that does xxx
 *
 * More detailed version (if necessary) - logs fatal for some reason
 */

func main()  {
	r := mux.NewRouter()
	routes.AreaRouter(r)
	http.Handle("/", r)

	log.Fatal(http.ListenAndServe("0.0.0.0:8080", r))

}

// func main()  {
// 	controllers.TriggerEachSecondes()
// }

