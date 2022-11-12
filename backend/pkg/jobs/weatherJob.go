/** @file weatherJob.go
 * @brief This file contain all the functions to handle the actions and reactions of the Weather API
 * @author Juliette Destang
 * @version
 */

// @cond

package jobs

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"AREA/pkg/utils"
)

// @endcond

/** @brief This function take a user id and activate his job on login
 * @param city string
 * @return float64 temperature
 */
func GetWeather(city string) (float64 ,error){

	url := "https://weather-by-api-ninjas.p.rapidapi.com/v1/weather?city=" + city

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", utils.GetEnv("RAPID_API_KEY"))
	req.Header.Add("X-RapidAPI-Host", utils.GetEnv("WEATHER_API_TOKEN"))

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

	temperature := weatherData["temp"]
 	// fmt.Println(string(body))

	// fmt.Println(weatherData)

	return temperature.(float64), nil
}

/** @brief An action that return true if the temperature is over N degrees
 * @param params string
 * @return bool
 */
func TemperatureIsOverN(params string) (bool) {
	paramsArr := utils.GetParams(params)
	fmt.Println(paramsArr)
	compareTemp, _ := strconv.ParseFloat(paramsArr[1], 64)
	temperature, weatherErr := GetWeather(paramsArr[0])
	fmt.Println(compareTemp)
	fmt.Println(temperature)
	if (weatherErr != nil) {
		fmt.Println(weatherErr)
		return false
	}

	if (temperature > compareTemp && temperature != 0) {
		return true
	} else { 
		return false
	}
}

/** @brief An action that return true if the temperature is under N degrees
 * @param params string
 * @return bool
 */
func TemperatureIsUnderrN(params string) (bool) {
	paramsArr := utils.GetParams(params)
	compareTemp, _ := strconv.ParseFloat(paramsArr[1], 64)
	temperature, weatherErr := GetWeather(paramsArr[0])
	if (weatherErr != nil) {
		fmt.Println(weatherErr)
		return false
	}

	if (temperature < compareTemp && temperature != 0) {
		return true
	} else { 
		return false
	}
}