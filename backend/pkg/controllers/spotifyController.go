package controllers

import (
	"net/http"
	"AREA/pkg/utils"
	"AREA/pkg/models"
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
)

func GetSpotifyUrl(w http.ResponseWriter, r *http.Request) {
	utils.EnableCors(&w)
	spotifyID := utils.GetEnv("SPOTIFY_ID");
	res, _ := json.Marshal(fmt.Sprintf("https://accounts.spotify.com/authorize?client_id=%s&redirect_uri=http://localhost:8080/spotify/auth&response_type=code&s&scope=user-modify-playback-state user-read-private user-read-currently-playing user-library-modify&state=random", spotifyID))
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

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

func AuthSpotify(w http.ResponseWriter, r *http.Request) {

	data := url.Values{}

	client := &http.Client{
	Timeout: time.Second * 10,
	}
	data.Set("client_id", utils.GetEnv("SPOTIFY_ID"))
	data.Set("client_secret", utils.GetEnv("SPOTIFY_SECRET"))
	data.Set("grant_type", "authorization_code")
	data.Set("redirect_uri", "http://localhost:8080/spotify/auth")
	data.Set("code", r.FormValue("code"))
	encodedData := data.Encode()
	
	const tokenurl = "https://accounts.spotify.com/api/token";

	req, err := http.NewRequest("POST", tokenurl, strings.NewReader(encodedData))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res, _ := json.Marshal("bad request")
		w.Write(res)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	response, _ := client.Do(req)

	body, _ := ioutil.ReadAll(response.Body)
	spotifyResponse := make(map[string]interface{})

	errorUnmarshal := json.Unmarshal(body, &spotifyResponse)
	if errorUnmarshal != nil {
	    log.Fatal(errorUnmarshal)
	}	

	accessToken := spotifyResponse["access_token"]
	refreshToken := spotifyResponse["refresh_token"]

	requestUser, _ := GetUser(w, r)


	models.SetUserToken(strconv.FormatUint(uint64(requestUser.ID), 10), "spotify_token", fmt.Sprintf("%s", accessToken))
	models.SetUserToken(strconv.FormatUint(uint64(requestUser.ID), 10), "spotify_refresh_token", fmt.Sprintf("%s", refreshToken))
	// song := GetSongByName(requestUser.ID ,"beatit")
	// PlayASong(requestUser.ID, song)
	// w.Write(spotifyResponse["access_token"])
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
	fmt.Println()
	fmt.Println(req)
	fmt.Println()

	response, _ := client.Do(req)

	body, _ := ioutil.ReadAll(response.Body)

	choosenSong := gjson.GetBytes(body, "tracks.items.0.uri")


	return choosenSong.String()
}

func PlayASong(userID uint, trackID string) {

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

func AddSongToQueue() {

}
