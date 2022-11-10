package controllers

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

	"AREA/pkg/utils"
	"AREA/pkg/models"
)

func GetSpotifyUrl(w http.ResponseWriter, r *http.Request) {
	utils.EnableCors(&w)
	spotifyID := utils.GetEnv("SPOTIFY_ID");
	res, _ := json.Marshal(fmt.Sprintf("https://accounts.spotify.com/authorize?client_id=%s&redirect_uri=http://localhost:8080/spotify/auth&response_type=code&s&scope=user-modify-playback-state user-read-private user-read-currently-playing user-library-modify&state=random", spotifyID))
	w.WriteHeader(http.StatusOK)
	w.Write(res)
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
	http.Redirect(w, r, "http://localhost:8081/user/services", http.StatusSeeOther)
	// song := GetSongByName(requestUser.ID ,"beatit")
	// PlayASong(requestUser.ID, song)
	// w.Write(spotifyResponse["access_token"])
}