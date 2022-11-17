package jobs

import (
	"AREA/pkg/utils"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/tidwall/gjson"
)

func GetMoonData() (string, error) {

	url := "https://moon-api1.p.rapidapi.com/distance?date-time=2009-07-11-09-30-00&timezone=%2B3&length-units=km"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", "54fb216729msh1db59bd41d901b7p12938ajsn6b6525d7a1c2")
	req.Header.Add("X-RapidAPI-Host", "moon-api1.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		myErr := errors.New("moon api down")
		return "", myErr
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	result := gjson.GetBytes(body, "distance").String()

	return result, nil
}

func MoonDistanceIsOverN(params string) bool {
	paramsArr := utils.GetParams(params)
	if len(paramsArr) != 1 {
		fmt.Fprintln(os.Stderr, paramsArr, "params passed are not correct")
		return false
	}

	compareTemp, _ := strconv.ParseFloat(paramsArr[0], 64)
	distance, weatherErr := GetMoonData()
	if weatherErr != nil {
		fmt.Fprintln(os.Stderr, weatherErr)
		return false
	}
	distanceValue, _ := strconv.ParseFloat(distance, 64)
	if distanceValue > compareTemp {
		return true
	} else {
		return false
	}
}

func MoonDistanceIsUnderN(params string) bool {
	paramsArr := utils.GetParams(params)
	if len(paramsArr) != 1 {
		fmt.Fprintln(os.Stderr, paramsArr, "params passed are not correct")
		return false
	}

	compareTemp, _ := strconv.ParseFloat(paramsArr[0], 64)
	distance, weatherErr := GetMoonData()
	if weatherErr != nil {
		fmt.Fprintln(os.Stderr, weatherErr)
		return false
	}
	distanceValue, _ := strconv.ParseFloat(distance, 64)
	if distanceValue < compareTemp {
		return true
	} else {
		return false
	}
}
