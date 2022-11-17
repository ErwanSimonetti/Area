/** @file weatherJob.go
 * @brief This file contain all the functions to handle the actions and reactions of the Weather API
 * @author Juliette Destang
 *
 */

// @cond

package jobs

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/tidwall/gjson"

	"AREA/pkg/utils"
)

// @endcond

/** @brief This function take a user id and activate his job on login
 * @param city string
 * @return float64 temperature
 */
func GetLovePercentage(name1 string, name2 string) (float64, error) {

	url := "https://love-calculator.p.rapidapi.com/getPercentage?sname=" + name1 + "&fname=" + name2

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", "54fb216729msh1db59bd41d901b7p12938ajsn6b6525d7a1c2")
	req.Header.Add("X-RapidAPI-Host", "love-calculator.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		myErr := errors.New("weather api is down")
		return 0, myErr
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	result := gjson.GetBytes(body, "percentage").String()

	loveValue, _ := strconv.ParseFloat(result, 64)

	return loveValue, nil
}

/** @brief An action that return true if the temperature is under N degrees
 * @param params string
 * @return bool
 */
func IsLoveHigherThan(params string) bool {
	paramsArr := utils.GetParams(params)
	if len(paramsArr) != 3 {
		fmt.Fprintln(os.Stderr, paramsArr, "params passed are not correct")
		return false
	}
	name1 := paramsArr[0]
	name2 := paramsArr[1]
	percentage, _ := strconv.ParseFloat(paramsArr[2], 64)

	loveValue, _ := GetLovePercentage(name1, name2)
	if loveValue >= percentage {
		return true
	}
	return false
}
