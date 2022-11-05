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
	"encoding/base64"
)

func GetSpotifyUrl(w http.ResponseWriter, r *http.Request) {
	utils.EnableCors(&w)
	spotifyID := utils.GetEnv("SPOTIFY_ID");
	res, _ := json.Marshal(fmt.Sprintf("https://accounts.spotify.com/authorize?client_id=%s&redirect_uri=http://localhost:8080/spotify/auth&response_type=code&s&scope=user-modify-playback-state user-read-private user-read-currently-playing user-library-modify&state=random", spotifyID))
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func RefreshSpotifyToken(refreshToken string) *(map[string]interface{}){

	client := &http.Client{
		Timeout: time.Second * 10,
	}

		refreshurl := "https://accounts.spotify.com/api/token"
		
		fmt.Println("refresh token HERE", refreshToken)
		
		refreshData := url.Values{}
		refreshData.Set("refresh_token", refreshToken)
		refreshData.Set("grant_type", "refresh_token")
		refreshData.Set("client_id", utils.GetEnv("SPOTIFY_ID"))
		refreshEncodedData := refreshData.Encode()

		refreshreq, err := http.NewRequest("POST", refreshurl, strings.NewReader(refreshEncodedData))
		if err != nil {
			// w.WriteHeader(http.StatusBadRequest)
			// res, _ := json.Marshal("bad request")
			// w.Write(res)
		}
		refreshreq.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		refreshreq.Header.Add("Authorization", "Basic " + base64.StdEncoding.EncodeToString([]byte((utils.GetEnv("SPOTIFY_ID") + ":" +utils.GetEnv("SPOTIFY_SECRET")))))

		refreshResponse, _ := client.Do(refreshreq)

		refreshBody, _ := ioutil.ReadAll(refreshResponse.Body)
		spotifyRefreshResponse := make(map[string]interface{})

		refresherrorUnmarshal := json.Unmarshal(refreshBody, &spotifyRefreshResponse)
		if refresherrorUnmarshal != nil {
			log.Fatal(refresherrorUnmarshal)
		}

		fmt.Println(spotifyRefreshResponse)
		return &spotifyRefreshResponse

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

	refreshToken := spotifyResponse["refresh_token"]
	spotifyResponse = *RefreshSpotifyToken(fmt.Sprintf("%s", refreshToken))
	//if (spotifyResponse["error"] != "") {
		
		//curl -X POST "https://example.com/v1/refresh" -H "Content-Type: application/x-www-form-urlencoded" --data "refresh_token=NgCXRK...MzYjw"
	//}

	fmt.Println(spotifyResponse)
	accessToken := spotifyResponse["access_token"]
	fmt.Println("access", accessToken)

	song := GetSongByName("minecraft", fmt.Sprintf("%s", accessToken))
	PlayASong(fmt.Sprintf("%s", accessToken), song)
	// w.Write(spotifyResponse["access_token"])
}

func GetSongByName(songName string, token string) (string){
	url := fmt.Sprintf("https://api.spotify.com/v1/search?q=%s&type=track", songName)

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	fmt.Println("TOKEN", token)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("can't get song name")
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer " + token)
	fmt.Println()
	fmt.Println(req)
	fmt.Println()

	response, _ := client.Do(req)
	fmt.Println(response)

	body, _ := ioutil.ReadAll(response.Body)
	// spotifyResponse := make(map[string]interface{})
	// fmt.Println("ok ici c bon")

	// errorUnmarshal := json.Unmarshal(body, &spotifyResponse)
	// if errorUnmarshal != nil {
	//     log.Fatal(errorUnmarshal)
	// }
	fmt.Println("ok ici c bon")
	choosenSong := gjson.GetBytes(body, "tracks.items.0.uri")
	fmt.Println(choosenSong)


	return fmt.Sprintf("%s",choosenSong)
}

func PlayASong(token string, trackID string) {

	url := fmt.Sprintf("https://api.spotify.com/v1/me/player/add-to-queue?uri=%s", trackID)

	fmt.Println(trackID)

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		log.Fatal("can't get song name")
	}
	req.Header.Add("Authorization", "Bearer " + token)

	response, _ := client.Do(req)
	response = response
	// fmt.Println("lalala")
	// body, _ := ioutil.ReadAll(response.Body)

	// spotifyResponse := make(map[string]interface{})

	// errorUnmarshal := json.Unmarshal(body, &spotifyResponse)
	// if errorUnmarshal != nil {
	//     log.Fatal(errorUnmarshal)
	// }
	fmt.Println("looooo")

	// fmt.Println(spotifyResponse)

}

func AddSongToQueue() {

}


// exports.addSongToQueue = async function(area)
// {
// 	var token = await global.findInDbAsync(global.CollectionToken, {user_id : area.user_id, service : global.Services.Spotify});
//     if (!token)
// 		return 'No access token provide';
// 	let track_id = await getSongByName(global.getParam(area.reaction.params, 'song_name'), token.access_token);
// 	if (!track_id)
// 		return 'Spotify didn\'t find the song';
// 	if (!(await addSongToQueue_request(track_id, token.access_token)))
// 		return 'Spotify failed to add song queue';
// };

// exports.playSong = async function (area)
// {
// 	var token = await global.findInDbAsync(global.CollectionToken, {user_id : area.user_id, service : global.Services.Spotify});
//     if (!token)
// 		return 'No access token provide';
// 	let track_id = await getSongByName(global.getParam(area.reaction.params, 'song_name'), token.access_token);
// 	if (!track_id)
// 		return 'Spotify didn\'t find the song';
// 	if (!(await addSongToQueue_request(track_id, token.access_token)))
// 		return 'Spotify failed add song queue';
// 	let url = 'https://api.spotify.com/v1/me/player/next';
//     let response = await fetch(url, {
//         'method': 'POST',
//         'headers' : {'Authorization' : 'Bearer ' + token.access_token}
// 	});
// 	if (response.status != 204)
// 		return 'Spotify failed to play next song';
// }