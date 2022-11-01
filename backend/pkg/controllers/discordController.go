package controllers

import (
	"fmt"
	"log"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
	"encoding/json"
	// "github.com/DisgoOrg/disgohook"
	"AREA/pkg/utils"
	"AREA/pkg/models"
)


var state = "random"

func AuthDiscord(w http.ResponseWriter, r *http.Request){

	authUrl := "https://discordapp.com/api/v6/oauth2/token";

	client := &http.Client{
		Timeout: time.Second * 10,
	}
	
	data := url.Values{}
	data.Set("client_id", utils.GetEnv("DISCORD_CLIENT_ID"))
	data.Set("client_secret", utils.GetEnv("DISCORD_CLIENT_SECRET"))
	data.Set("grant_type", "authorization_code")
	data.Set("redirect_uri", "http://localhost:8080/discord/auth")
	data.Set("scope", "webhook.incoming")
	data.Set("code", r.FormValue("code"))
	encodedData := data.Encode()
	fmt.Println(encodedData)
	req, err := http.NewRequest("POST", authUrl, strings.NewReader(encodedData))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res, _ := json.Marshal("bad request")
		w.Write(res)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	response, _ := client.Do(req)

	body, _ := ioutil.ReadAll(response.Body)
	jsonWebhook := make(map[string]interface{})
	errorUnmarshal := json.Unmarshal(body, &jsonWebhook)
	if errorUnmarshal != nil {
	    log.Fatal(errorUnmarshal)
	}

	cookieValue, cookieErr := r.Cookie("userID")
	if cookieErr != nil {
		panic(err.Error())
	}

	fmt.Println(jsonWebhook["webhook"])
	address := jsonWebhook["webhook"].(map[string]interface{})

	webhookId := fmt.Sprintf("%s", address["id"])
	webhookToken := fmt.Sprintf("%s", address["token"])

	models.SetUserToken(cookieValue.Value, "discord_id", webhookId)
	models.SetUserToken(cookieValue.Value, "discord_token", webhookToken)


	// webhook, err := disgohook.NewWebhookClientByToken(nil, nil, messageUrl)
	// msg := "test of Dana's tribe"

	// Imessage, err := webhook.SendContent(msg)

	// Imessage = Imessage
}

func SendMessage(userID int) {

	// userToken := *models.FindUserToken(userID)

    // messageUrl := fmt.Sprintf("%s/%s", userToken.DiscordId, userToken.DiscordToken)

    // webhook, _ := disgohook.NewWebhookClientByToken(nil, nil, messageUrl)
    // msg := "test of Dana's tribe"

    // Imessage, _ := webhook.SendContent(msg)
}

func GetDiscordUrl(w http.ResponseWriter, r *http.Request) {
	utils.EnableCors(&w)
	discordID := utils.GetEnv("DISCORD_CLIENT_ID");
	res, _ := json.Marshal(fmt.Sprintf("https://discord.com/api/oauth2/authorize?client_id=%s&redirect_uri=http://localhost:8080/discord/auth&response_type=code&scope=webhook.incoming&permissions=536870912", discordID))
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}	