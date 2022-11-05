package controllers

// import (
// 	"context"
// 	"fmt"
// 	"github.com/zmb3/spotify/v2/auth"
// 	"log"
// 	"net/http"
// 	"encoding/json"
// 	"github.com/zmb3/spotify/v2"
// )

// func GetSpotifyUrl(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("ok")
// 	url := auth.AuthURL(state)
// 	res, _ := json.Marshal(url)
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)

// }

// // redirectURI is the OAuth redirect URI for the application.
// // You must register an application at Spotify's developer portal
// // and enter this value.
// const redirectURI = "http://localhost:8080/spotify/auth"

// var (
// 	auth  = spotifyauth.New(spotifyauth.WithRedirectURL(redirectURI), spotifyauth.WithScopes(spotifyauth.ScopeUserReadCurrentlyPlaying, spotifyauth.ScopeUserReadPlaybackState, spotifyauth.ScopeUserModifyPlaybackState))
// 	ch    = make(chan *spotify.Client)
// )

// func main() {
// 	// first start an HTTP server
// 	// http.HandleFunc("/callback", completeAuth)
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		log.Println("Got request for:", r.URL.String())
// 	})
// 	go func() {
// 		err := http.ListenAndServe(":8080", nil)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 	}()

// 	url := auth.AuthURL(state)
// 	fmt.Println("Please log in to Spotify by visiting the following page in your browser:", url)

// 	// wait for auth to complete
// 	client := <-ch

// 	// use the client to make calls that require authorization
// 	user, err := client.CurrentUser(context.Background())
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("You are logged in as:", user.ID)
// }

// func AuthSpotify(w http.ResponseWriter, r *http.Request) {
// 	tok, err := auth.Token(r.Context(), state, r)
// 	if err != nil {
// 		http.Error(w, "Couldn't get token", http.StatusForbidden)
// 		log.Fatal(err)
// 	}
// 	if st := r.FormValue("state"); st != state {
// 		http.NotFound(w, r)
// 		log.Fatalf("State mismatch: %s != %s\n", st, state)
// 	}


// 	client := spotify.New(auth.Client(r.Context(), tok))
// 	// infinite loop ^^^

// 	w.Header().Set("Content-Type", "text/html")
// 	fmt.Fprintf(w, "Login Completed!")
// 	ch <- client
// }
