/** @file deezerController.go
 * @brief Oauth and reactions for Deezer
 * @author Erwan
 */

// @cond

package controllers

import (
	"fmt"
	"log"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"github.com/tidwall/gjson"
	"strconv"
	"strings"
	"encoding/json"
	"AREA/pkg/utils"
	"AREA/pkg/models"
	"html"
)

// @endcond


/** @brief gets a track, album, artist or label id through a search on Deezer's API
 *
 * @param[string] query
 * @param[string] queryType
 * @param[string] access_token
 *
 * @return queryId
 */
func GetQueryId(query string, queryType string, access_token string) (int64) {
	Url := "https://api.deezer.com/search?q=" + queryType + ":\"" + url.QueryEscape(query) +"\"&access_token=" + access_token
	resp, err := http.Get(Url)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("client: could not read response body: %s\n", err)
		os.Exit(1)
	}
	parsedResponse := make(map[string]interface{})

	errorUnmarshal := json.Unmarshal(body, &parsedResponse)
	if errorUnmarshal != nil {
	    log.Fatal(errorUnmarshal)
	}
	id := gjson.GetBytes(body, "data.0.id")
	idConverted := id.Int()
	return (idConverted)
}


/** @brief gets a playlist ID based on its EXACT name
 *
 * @param[string] access_token
 * @param[string] playlistName
 *
 * @return playlistId
 */
func GetPlaylistIdByName(access_token string, playlistName string) (int64) {
	Url := "https://api.deezer.com/user/me/playlists?limit=999&access_token=" + access_token
	log.Println("url = " + Url)
	resp, err := http.Get(Url)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("client: could not read response body: %s\n", err)
		os.Exit(1)
	}
	title := gjson.GetBytes(body, "data.0.title")
	for i := 0; i < 999; i++ {
		currentIndex := strconv.Itoa(i)
		title = gjson.GetBytes(body, "data." + currentIndex + ".title")
		if title.Str == playlistName {
			playlistId := gjson.GetBytes(body, "data." + currentIndex + ".id")
			idConverted := playlistId.Int()
			return (idConverted)
		}
	}
	return (-1)
}

/** @brief Adds a song to a playlist
 *
 * @param[string] access_token
 * @param[string] playlistId
 * @param[string] trackId
 */
func AddSongToPlaylist(access_token string, playlistId string, trackId string) {
	Url := "https://api.deezer.com/playlist/" + playlistId + "/tracks?songs=" + trackId + "&access_token=" + access_token
	log.Println("url = ")
	log.Println(Url)
	req, _ := http.NewRequest("POST", Url, nil)

	log.Println("request =")
	log.Println(req)
	hc := &http.Client{}
	resp, _ := hc.Do(req)

	log.Println("resp =")
	log.Println(resp)
}

/** @brief Returns the current user's ID
 *
 * @param[string] access_token
 *
 * @return userId
 */
func GetUserId(access_token string) (int64) {
	log.Println("access token: ")
	log.Println(access_token)
	Url := "https://api.deezer.com/user/me?access_token=" + access_token
	resp, err := http.Get(Url)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("client: could not read response body: %s\n", err)
		os.Exit(1)
	}
	id := gjson.GetBytes(body, "id")
	idConverted := id.Int()
	return (idConverted)
}

/** @brief Connects the user and does some reactions
 *
 * Connexion via OAuth2 and then researching a playlist and a song,
 * adding the song to the playlist
 */
func AuthDeezer(w http.ResponseWriter, r *http.Request) {
    keys, ok := r.URL.Query()["code"]
    
    if !ok || len(keys[0]) < 1 {
        log.Fatalln("Url Param 'code' is missing")
        os.Exit(1)
    }
    code := keys[0]

    log.Println("Url Param 'code' is: " + string(code))

	authUrl := "https://connect.deezer.com/oauth/access_token.php"
	data := url.Values{}
	data.Set("app_id", "564742")
	data.Set("secret", "7fd8db961fb009848532ca1c901bbc76")
	data.Set("code", string(code))
	encodedData := data.Encode()
	newUrl := authUrl + "?" + encodedData

	resp, err := http.Get(newUrl)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("client: could not read response body: %s\n", err)
		os.Exit(1)
	}
	access_token := strings.Split(string(body), "=")[1]
	access_token = strings.Split(access_token, "&")[0]
	app_id, _ := strconv.Atoi(utils.GetEnv("DEEZER_APP_ID"))
	models.SetUserToken(strconv.FormatUint(uint64(app_id), 10), "deezer_token", access_token)
	userId := GetUserId(access_token)
	trackId := GetQueryId("Thriller", "track", access_token)
	playlistId := GetPlaylistIdByName(access_token, "API test")
	log.Println("user id =")
	log.Println(userId)
	AddSongToPlaylist(access_token, strconv.FormatInt(playlistId, 10), strconv.FormatInt(trackId, 10))
}

/** @brief Returns the URL for OAuth2
 *
 * It has the full URL to ask for permissions and the redirection URL asked by the frontend
 */
func GetDeezerUrl(w http.ResponseWriter, r *http.Request) {
	utils.EnableCors(&w)
	deezerID := utils.GetEnv("DEEZER_APP_ID");
	res, _ := json.Marshal(html.EscapeString("https://connect.deezer.com/oauth/auth.php?app_id=" + deezerID + "&redirect_uri=http://localhost:8080/deezer/auth&perms=basic_access,email,offline_access,manage_library,listening_history"))
	log.Println("res = ")
	log.Println(string(res))
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "https://connect.deezer.com/oauth/auth.php?app_id=%s&redirect_uri=http://localhost:8080/deezer/auth&perms=basic_access,email,offline_access,manage_library,listening_history", deezerID)
}	
