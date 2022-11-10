package jobs

import (
	"net/http"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"strconv"
	"io/ioutil"
	"log"
	"time"
	"github.com/tidwall/gjson"
	"encoding/base64"

	"AREA/pkg/utils"
	"AREA/pkg/models"
)

func RefreshSpotifyToken(userID uint) {

	userToken := *models.FindUserToken(userID)
	client := &http.Client{
		Timeout: time.Second * 10,
	}

		refreshurl := "https://accounts.spotify.com/api/token"
		
		
		refreshData := url.Values{}
		refreshData.Set("refresh_token", userToken.SpotifyRefreshToken)
		refreshData.Set("grant_type", "refresh_token")
		refreshData.Set("client_id", utils.GetEnv("SPOTIFY_ID"))
		refreshEncodedData := refreshData.Encode()

		refreshreq, _ := http.NewRequest("POST", refreshurl, strings.NewReader(refreshEncodedData))
		// if err != nil {
			// w.WriteHeader(http.StatusBadRequest)
			// res, _ := json.Marshal("bad request")
			// w.Write(res)
		// }
		refreshreq.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		refreshreq.Header.Add("Authorization", "Basic " + base64.StdEncoding.EncodeToString([]byte((utils.GetEnv("SPOTIFY_ID") + ":" +utils.GetEnv("SPOTIFY_SECRET")))))

		refreshResponse, _ := client.Do(refreshreq)

		refreshBody, _ := ioutil.ReadAll(refreshResponse.Body)

		spotifyRefreshResponse := make(map[string]interface{})

		refresherrorUnmarshal := json.Unmarshal(refreshBody, &spotifyRefreshResponse)
		if refresherrorUnmarshal != nil {
			log.Fatal(refresherrorUnmarshal)
		}
		accessToken := spotifyRefreshResponse["access_token"]

		models.SetUserToken(strconv.FormatUint(uint64(userID), 10), "spotify_token", fmt.Sprintf("%s", accessToken))
}

func GetSongByName(userID uint , songName string) (string){

	userToken := *models.FindUserToken(userID)

	url := fmt.Sprintf("https://api.spotify.com/v1/search?q=%s&type=track", songName)

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	fmt.Println("TOKEN", userToken.SpotifyToken)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		RefreshSpotifyToken(userID)
		log.Fatal("can't get song name")
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer " + userToken.SpotifyToken)

	response, _ := client.Do(req)

	body, _ := ioutil.ReadAll(response.Body)

	choosenSong := gjson.GetBytes(body, "tracks.items.0.uri")


	return choosenSong.String()
}

func AddSongToQueue(userID uint, params string) {
	trackID := GetSongByName(userID, params)
	userToken := *models.FindUserToken(userID)

	url := fmt.Sprintf("https://api.spotify.com/v1/me/player/add-to-queue?uri=%s", trackID)

	client := &http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		log.Fatal("can't get song name")
	}
	req.Header.Add("Authorization", "Bearer " + userToken.SpotifyToken)

	client.Do(req)
	// response = response

}