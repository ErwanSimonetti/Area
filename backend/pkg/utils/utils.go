/** @file utils.go
 * @brief This file contains all the functions to handle mondain actions
 * @author Juliette Destang
 * @version
 */

//cond
package utils

import(
	"encoding/json"
	"io/ioutil"
	"net/http"
	"github.com/joho/godotenv"
	"os"
	"strings"
	"log"
)

//endcond

/** @brief Returns true if the string is contained by the array of string
 * @param s []string, str string
 * @return bool
 */
func ArrayContainsString(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

/** @brief Splits a string by '@@@' and returns it
 * @param params string
 * @return []string
 */
func GetParams(params string) []string {
	split := strings.Split(params, "@@@")
	return split
}

/** @brief Gets all the information contained in the body of the request
 * @param r *http.Request, x interface{}
 */
func ParseBody(r *http.Request, x interface{}){
	if body, err := ioutil.ReadAll(r.Body); err == nil{
		if err := json.Unmarshal([]byte(body), x); err != nil{
			return 
		}
	}
}

/** @brief Enables the Cors so the front can access the response
 * @param w *http.ResponseWriter
 */
func EnableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:8081")
}

/** @brief Loads the env from our .env
 * @param key string
 * @return string
 */
func GetEnv(key string) string {
	err := godotenv.Load(".env")
  
	if err != nil {
	  log.Fatal("Error loading .env file")
	}
  
	return os.Getenv(key)
}