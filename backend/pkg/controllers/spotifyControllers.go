package controllers

import (
	"fmt"
	"context"
	"log"
	"io/ioutil"
	"net/http"
	"net/url"
	
	"golang.org/x/oauth2"
	"github.com/ravener/discord-oauth2"
	"github.com/zmb3/spotify"
	"github.com/jasonlvhit/gocron"
	"golang.org/x/oauth2/clientcredentials"
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

	conf := &oauth2.Config{
		RedirectURL: "http://localhost:8080/auth/discord",
		ClientID: "1033382176785432656",
		ClientSecret: "XYlj-YTnAqQ-7HNes4I5xRTIdIKdLvyZ",
		Scopes: []string{discord.ScopeActivitiesRead, discord.ScopeIdentify, discord.ScopeMessagesRead},
		Endpoint: discord.Endpoint,
	}


		data := url.Values{
			"content":       {"John Doe"},
			"tts": {"false"},
		}


		// byts, _ := json.Marshal(payload)	
	token, _ := conf.Exchange(context.Background(), r.FormValue("code"))
	fmt.Println(token)

	// res, err := conf.Client(context.Background(), token).Get("https://discord.com/api/users/@me")
	message, err := conf.Client(context.Background(), token).Get("https://discord.com/api/channels/1021410887191507004/messages")

	ecrire, err := conf.Client(context.Background(), token).PostForm("https://discord.com/api/channels/1021410887191507004/messages", data)
	ecrire = ecrire
	fmt.Println("message;")
	fmt.Println(message)
	// byts = byts
	if err != nil || message.StatusCode != 200 {
		w.WriteHeader(http.StatusInternalServerError)
		if err != nil {
			w.Write([]byte(err.Error()))
		} else {
			w.Write([]byte(message.Status))
		}
		return
	}

	defer message.Body.Close()

	body, err := ioutil.ReadAll(message.Body)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(body)

}