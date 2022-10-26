package controllers

import (
	// "fmt"
	// "context"
	"log"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	// "strconv"
	"strings"
	// "time"
	// "github.com/jasonlvhit/gocron"
	// "golang.org/x/oauth2/clientcredentials"
	// "encoding/json"
	// "github.com/DisgoOrg/disgohook"
)

func GetUserInfos(access_token string) {
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
	log.Println("body =")
	log.Println(string(body))
}

func DeezerTest(w http.ResponseWriter, r *http.Request) {
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
	GetUserInfos(access_token)
}