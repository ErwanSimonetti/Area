package controllers

import (
	"net/http"
	"AREA/pkg/utils"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"strconv"
	"io/ioutil"
	"log"
	"time"
	"github.com/tidwall/gjson"
)

func GetSpotifyUrl(w http.ResponseWriter, r *http.Request) {
	utils.EnableCors(&w)
	spotifyID := utils.GetEnv("SPOTIFY_CLIENT_ID");
	res, _ := json.Marshal(fmt.Sprintf("https://accounts.spotify.com/authorize?client_id=%s&redirect_uri=http://localhost:8080/spotify/auth&response_type=code&s&scope=user-modify-playback-state user-read-private user-read-currently-playing user-library-modify", spotifyID))
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func AuthSpotify(w http.ResponseWriter, r *http.Request) {

	data := url.Values{}

	client := &http.Client{
	Timeout: time.Second * 10,
	}
	data.Set("client_id", utils.GetEnv("SPOTIFY_CLIENT_ID"))
	data.Set("client_secret", utils.GetEnv("SPOTIFY_CLIENT_SECRET"))
	data.Set("grant_type", "authorization_code")
	data.Set("redirect_uri", "http://localhost:8080/spotify/auth")
	data.Set("code", r.FormValue("code"))
	encodedData := data.Encode()
	
	const url = "https://accounts.spotify.com/api/token";

	req, err := http.NewRequest("POST", url, strings.NewReader(encodedData))
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
	fmt.Println(accessToken)

	getSongByName("thriller", fmt.Sprintf("%s", accessToken))
	// w.Write(spotifyResponse["access_token"])
}

func getSongByName(songName string, token string) {
	url := fmt.Sprintf("https://api.spotify.com/v1/search?q=%s&type=track", songName)

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("can't get song name")
	}
	req.Header.Add("Authorization", "Bearer " + token)

	response, _ := client.Do(req)

	
	body, _ := ioutil.ReadAll(response.Body)
	choosenSong := gjson.GetBytes(body, "tracks.items.0.uri")

	choosenSong = choosenSong
}
// async function getSongByName(song_name, token)
// {
//     let url = 'https://api.spotify.com/v1/search?q=' + song_name + '&type=track'
//     let response = await fetch(url, {
//         'method': 'GET',
//         'headers' : { 'Authorization' : 'Bearer ' + token }
//     });
//     if (response.status != 200)
// 		return false;
// 	let resjson = await response.json();
// 	var track_id;
// 	try {
// 		track_id = resjson.tracks.items[0].uri;
// 	} catch (err) {
// 		return null;
// 	}
// 	return track_id;
// }


// func AuthDiscord(w http.ResponseWriter, r *http.Request){

// 	authUrl := "https://discordapp.com/api/v6/oauth2/token";

// 	client := &http.Client{
// 		Timeout: time.Second * 10,
// 	}
	
// 	data := url.Values{}
// 	data.Set("client_id", utils.GetEnv("DISCORD_CLIENT_ID"))
// 	data.Set("client_secret", utils.GetEnv("DISCORD_CLIENT_SECRET"))
// 	data.Set("grant_type", "authorization_code")
// 	data.Set("redirect_uri", "http://localhost:8080/discord/auth")
// 	data.Set("scope", "webhook.incoming")
// 	data.Set("code", r.FormValue("code"))
// 	encodedData := data.Encode()

// 	req, err := http.NewRequest("POST", authUrl, strings.NewReader(encodedData))
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		res, _ := json.Marshal("bad request")
// 		w.Write(res)
// 	}
// 	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
// 	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
// 	response, _ := client.Do(req)
// 	fmt.Println("ici")

// 	body, _ := ioutil.ReadAll(response.Body)
// 	jsonWebhook := make(map[string]interface{})
// 	fmt.Println(body)
// 	errorUnmarshal := json.Unmarshal(body, &jsonWebhook)
// 	if errorUnmarshal != nil {
// 	    log.Fatal(errorUnmarshal)
// 	}

// 	requestUser, err := GetUser(w, r)

// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		res, _ := json.Marshal("bad request")
// 		w.Write(res)
// 	}

// 	fmt.Println(jsonWebhook["webhook"])
// 	address := jsonWebhook["webhook"].(map[string]interface{})

// 	webhookId := fmt.Sprintf("%s", address["id"])
// 	webhookToken := fmt.Sprintf("%s", address["token"])

// 	models.SetUserToken(strconv.FormatUint(uint64(requestUser.ID), 10), "discord_id", webhookId)
// 	models.SetUserToken(strconv.FormatUint(uint64(requestUser.ID), 10), "discord_token", webhookToken)
// }