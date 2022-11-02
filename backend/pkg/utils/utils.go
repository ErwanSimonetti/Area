package utils

import(
	"encoding/json"
	"io/ioutil"
	"net/http"
	"github.com/joho/godotenv"
	"os"
	"log"
)

func ParseBody(r *http.Request, x interface{}){
	if body, err := ioutil.ReadAll(r.Body); err == nil{
		if err := json.Unmarshal([]byte(body), x); err != nil{
			return 
		}
	}
}

func EnableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func GetEnv(key string) string {
	err := godotenv.Load(".env")
  
	if err != nil {
	  log.Fatal("Error loading .env file")
	}
  
	return os.Getenv(key)
  }