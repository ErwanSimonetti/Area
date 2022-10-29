package controllers

import (
	"log"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"AREA/pkg/utils"
)

var temperature float64 = 0

func GetWeather(w http.ResponseWriter, r *http.Request) {

	url := "https://open-weather13.p.rapidapi.com/city/paris"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", utils.GetEnv("RAPID_API_KEY"))
	req.Header.Add("X-RapidAPI-Host", "open-weather13.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res, _ := json.Marshal("bad request")
		w.Write(res)
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	weatherData := make(map[string]interface{})
	errorUnmarshal := json.Unmarshal(body, &weatherData)
	if errorUnmarshal != nil {
	    log.Fatal(errorUnmarshal)
	}

	address := weatherData["main"].(map[string]interface{})

	temperature, _ = address["temp"].(float64)

}

func TemperatureIsUnder24() bool {
	if (temperature < 24.0 && temperature != 0) {
		return true
	} else { 
		return false
	}
}

func TemperatureIsOver24() bool {
	if (temperature > 24.0 && temperature != 0) {
		return true
	} else { 
		return false
	}
}