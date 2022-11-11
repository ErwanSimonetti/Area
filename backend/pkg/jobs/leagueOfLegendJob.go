/** @file leagueOfLeagendJob.go
 * @brief This file contain all the functions to handle the actions and reactions of the Covid API
 * @author Juliette Destang
 * @version
 */

// @conv
package jobs

import (
	"fmt"
	"net/http"
	"io/ioutil"
	// "encoding/json"
	"strconv"
	"github.com/tidwall/gjson"
	"errors"
	"AREA/pkg/utils"
	// "log"
	"strings"
)

// @endconv

/** @brief Retrieves all the data concerning a given player from the League of legend API
 * @param playerName string
 * @return []byte ,error
 */
func GetLeagueStat(playerName string)([]byte ,error) {

	url := "https://lol_stats.p.rapidapi.com/euw1/" + playerName

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", utils.GetEnv("RAPID_API_KEY"))
	req.Header.Add("X-RapidAPI-Host", utils.GetEnv("LEAGUE_API_TOKEN"))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		myErr := errors.New("league api")
		return nil, myErr
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	
	return body, nil

}

/** @brief Returns true if the player played teemo recently
 * @param params string
 * @return bool
 */
func IsPlayingTeemo(params string) (bool) {
	leagueData, Err := GetLeagueStat(params)
	if (Err != nil) {
		fmt.Println(Err)
		return false
	}
	maybeTeemo := gjson.GetBytes(leagueData, "mostPlayedChamps.0.champName")
	if (maybeTeemo.String() == "Teemo") {
		return true
	} else {
		return false
	}
}

/** @brief Returns true if the player played winrate is over a given value
 * @param params string
 * @return bool
 */
func WinrateIsOverN(params string) (bool) {
	leagueData, Err := GetLeagueStat(params)
	if (Err != nil) {
		fmt.Println(Err)
		return false
	}
	winrate := gjson.GetBytes(leagueData, "mostPlayedChamps.0.winrate")
	cleanWinrate := strings.TrimSuffix(winrate.String(), "%")
	floatWinrate, _ := strconv.ParseFloat(cleanWinrate, 64)
	if (floatWinrate > 50) {
		return true
	} else {
		return false
	}
}

/** @brief Returns true if the player played KDA is over a given value
 * @param params string
 * @return bool
 */
func KDAIsOverN(params string) (bool) {
	leagueData, Err := GetLeagueStat(params)
	if (Err != nil) {
		fmt.Println(Err)
		return false
	}
	KDA := gjson.GetBytes(leagueData, "mostPlayedChamps.0.kda")
	floatKDA, _ := strconv.ParseFloat(KDA.String(), 64)
	if (floatKDA > 3.0) {
		return true
	} else {
		return false
	}
}