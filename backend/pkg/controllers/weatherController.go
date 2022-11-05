package controllers

import (
	"log"
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"AREA/pkg/utils"
	"errors"
)


func GetWeather() (float64 ,error){

	url := "https://weather-by-api-ninjas.p.rapidapi.com/v1/weather?city=Paris"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", utils.GetEnv("WEATHER_API_KEY"))
	req.Header.Add("X-RapidAPI-Host", "weather-by-api-ninjas.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		myErr := errors.New("Weather api is down")
		return 0, myErr
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	weatherData := make(map[string]interface{})
	errorUnmarshal := json.Unmarshal(body, &weatherData)
	if errorUnmarshal != nil {
	    log.Fatal(errorUnmarshal)
	}

	temperature := weatherData["temp"]
 	fmt.Println(string(body))

	// fmt.Println(weatherData)

	return temperature.(float64), nil
}


func TemperatureIsOver24() (bool, error) {
	temperature, weatherErr := GetWeather()
	if (weatherErr != nil) {
		fmt.Println(weatherErr)
		return false, weatherErr
	}

	if (temperature > 24.0 && temperature != 0) {
		return true, nil
	} else { 
		return false, nil
	}
}