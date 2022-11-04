package controllers

import (
	"log"
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"errors"

	"AREA/pkg/utils"
)

func GetWeather() (float64, error){

	url := "https://open-weather13.p.rapidapi.com/city/paris"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", utils.GetEnv("RAPID_API_KEY"))
	req.Header.Add("X-RapidAPI-Host", "open-weather13.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		myErr := errors.New("weather api is down")
		return 0, myErr
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	weatherData := make(map[string]interface{})
	errorUnmarshal := json.Unmarshal(body, &weatherData)
	if errorUnmarshal != nil {
	    log.Fatal(errorUnmarshal)
	}
	fmt.Println(weatherData)
	// address := weatherData["main"].(map[string]interface{})

	// newTemperature, _ := address["temp"].(float64)

	// temperature := (newTemperature - 32) * 5/9

	return 8, nil
}

func TemperatureIsUnder24() (bool, error) {
	temperature, weatherErr := GetWeather()
	if (weatherErr != nil) {
		fmt.Println(weatherErr)
		return false, weatherErr
	}

	if (temperature < 24.0 && temperature != 0) {
		return true, nil
	} else { 
		return false, nil
	}
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