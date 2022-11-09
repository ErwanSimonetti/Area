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

func GetLeagueStat()([]byte ,error) {

	url := "https://lol_stats.p.rapidapi.com/euw1/0Pixelle0"

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

func IsPlayingTeemo() (bool) {
	leagueData, Err := GetLeagueStat()
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

func WinrateIsOver50() (bool) {
	leagueData, Err := GetLeagueStat()
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

func KDAIsOver3() (bool) {
	leagueData, Err := GetLeagueStat()
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