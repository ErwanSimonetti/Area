package controllers

import (
	"fmt"
	"context"
	"log"
	"io/ioutil"
	"net/http"
	// "bytes"
	"net/url"
	"strconv"
	"strings"
	"time"
	// "AREA/pkg/utils"
	// "github.com/gorilla/mux"
	
	// "golang.org/x/oauth2"
	// "github.com/ravener/discord-oauth2"
	"github.com/zmb3/spotify"
	"github.com/jasonlvhit/gocron"
	"golang.org/x/oauth2/clientcredentials"
	// "github.com/gin-contrib/cors"
	"encoding/json"
    // "github.com/gin-gonic/gin"
	"github.com/DisgoOrg/disgohook"
)

func ConnectionSpotify() {
	authConfig := &clientcredentials.Config{
		ClientID: "df401fc1c764474ebb8406b6751fe0f3",
		ClientSecret: "798834ab9a4a462bad2c994cb9a3b8e1",
		TokenURL: spotify.TokenURL,
	}

	accessToken, err := authConfig.Token(context.Background())

	if err != nil {
		log.Fatalf("error retrieve access token: %v", err)
	}

	client := spotify.Authenticator{}.NewClient(accessToken)

	playlistID := spotify.ID("37i9dQZF1DXcBWIGoYBM5M")
	playlist, err := client.GetPlaylist(playlistID)
	if err != nil {
		log.Fatalf("error retrieve playlist data: %v", err)
	}

	log.Println("playlist id:", playlist.ID)
	log.Println("playlist name:", playlist.Name)
	log.Println("playlist description:", playlist.Description)
}

// func myPrint() {
// 	fmt.Println("jaj")
// }

// var c chan int
// func handle(int) {
// }

// func TriggerEachSecondes() {
// 	select {
// 	// case m := <-c:
// 	// 	handle(m)
// 	case <-time.After(1 * time.Second):
// 		fmt.Println("timed out")
// 	}
// }

func task() {
	resp, err := http.Get("https://blog.logrocket.com/making-http-requests-in-go/")
	if err != nil {
	   log.Fatalln(err)
	}
 //We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
	   log.Fatalln(err)
	}
 //Convert the body to type string
	sb := string(body)
	log.Printf(sb)
}

func TriggerEachSecondes() {
    s := gocron.NewScheduler()
    s.Every(2).Second().Do(task)
    <- s.Start()
}

var state = "random"

func HelloDiscord(w http.ResponseWriter, r *http.Request){

	// conf := &oauth2.Config{
	// 	RedirectURL: "http://localhost:8080/auth/discord",
	// 	ClientID: "1033382176785432656",
	// 	ClientSecret: "XYlj-YTnAqQ-7HNes4I5xRTIdIKdLvyZ",
	// 	Scopes: []string{discord.ScopeWebhookIncoming },
	// 	Endpoint: discord.Endpoint,
	// }
	// // token, _ := conf.Exchange(context.Background(), r.FormValue("code"))
	// // fmt.Println(token)
	// toekn = token
	myurl := "https://discordapp.com/api/v6/oauth2/token";

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	data := url.Values{}
	data.Set("client_id", "1033382176785432656")
	data.Set("client_secret", "XYlj-YTnAqQ-7HNes4I5xRTIdIKdLvyZ")
	data.Set("grant_type", "authorization_code")
	data.Set("redirect_uri", "http://localhost:8080/auth/discord")
	data.Set("scope", "webhook.incoming")
	data.Set("code", r.FormValue("code"))
	encodedData := data.Encode()
	fmt.Println(encodedData)
	req, err := http.NewRequest("POST", myurl, strings.NewReader(encodedData))
	if err != nil {
		// return fmt.Errorf("Got error %s", err.Error())
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	response, _ := client.Do(req)

	body, _ := ioutil.ReadAll(response.Body)
	  
	m := make(map[string]interface{})
	jajerr := json.Unmarshal(body, &m)
	if jajerr != nil {
	    log.Fatal(jajerr)
	}
	address := m["webhook"].(map[string]interface{})
	messageUrl := fmt.Sprintf("%s/%s", address["id"], address["token"])
	webhook, err := disgohook.NewWebhookClientByToken(nil, nil, messageUrl)

	Imessage, err := webhook.SendContent("premiere action bande de SCHLAGUE QUI A LA PLUS GROSSE BITE MAINTENANT HEEEEEIN @everyone\nDe la part du bot: Timothée et JUliette on des grosses quequetes")
	webhook.SendContent("premiere action bande de SCHLAGUE QUI A LA PLUS GROSSE BITE MAINTENANT HEEEEEIN @everyone\nDe la part du bot: Timothée et JUliette on des grosses quequetes")

	Imessage = Imessage
	// Imessage, err := webhook.SendEmbeds(api.NewEmbedBuilder().
	// SetDescription("hello world!").
	// Build(),
	// )
	// Imessage, err := webhook.SendMessage(api.NewWebhookMessageCreateBuilder().
	// 	SetContent("hello world!").
	// 	Build(),
	// )

	// address := m["webhook"].(map[string]interface{})
	// messageUrl := fmt.Sprintf("https://discordapp.com/api/webhooks/%s/%s", address["id"], address["token"])
	// // const messageUrl = "https://discordapp.com/api/webhooks/" + address["id"] + "/"{address["token"];

	// messageData := url.Values{}
	// messageData.Add("content", "je suce")
	// // messageData.Add("username", "salut lel")
	// // messageData.Add("avatar", "")
	// messageEncodedData := messageData.Encode()
	// req, _ = http.NewRequest("POST", messageUrl, strings.NewReader(messageEncodedData))
	// if err != nil {
	// 	// return fmt.Errorf("Got error %s", err.Error())
	// }
	// req.Header.Add("Content-Type", "application/json")
	// req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	// newresponse, _ := client.Do(req)

	// fmt.Println(messageUrl)
	// fmt.Println(newresponse)
	// let body = {
	// 	"content" : global.getParam(area.reaction.params, "message"),
	// 	username : global.getParam(area.reaction.params, "username"),
	// 	avatar_url : global.getParam(area.reaction.params, "avatar")
	// };
	// let response = await fetch(url, {
	// 	"method": "POST",
	// 	"body" : JSON.stringify(body),
	// 	"headers" : {"Content-Type" : "application/json"}
	// });
	// res, err := conf.Client(context.Background(), token).Get("https://discord.com/api/users/@me")


	// message, err := conf.Client(context.Background(), token).Get("https://discord.com/api/channels/1021410887191507004/messages")

	// byts = byts
	// if err != nil || res.StatusCode != 200 {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	if err != nil {
	// 		w.Write([]byte(err.Error()))
	// 	} else {
	// 		w.Write([]byte(res.Status))
	// 	}
	// 	return
	// }

	// defer res.Body.Close()

	// body, err := ioutil.ReadAll(res.Body)

	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	w.Write([]byte(err.Error()))
	// 	return
	// }

	// w.Write(body)

}