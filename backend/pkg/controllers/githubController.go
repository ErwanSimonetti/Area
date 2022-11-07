package controllers

import (
	"net/http"
	"AREA/pkg/utils"
	"AREA/pkg/models"
	"encoding/json"
	"strconv"
	"fmt"
	"time"
	"github.com/tidwall/gjson"
	"io/ioutil"
	// "log"
	"bytes"
)

func GetGithubUrl(w http.ResponseWriter, r *http.Request) {
	utils.EnableCors(&w)
	githubID := utils.GetEnv("GITHUB_ID");
	res, _ := json.Marshal(fmt.Sprintf("https://github.com/login/oauth/authorize?client_id=%s&scope=admin:repo_hook repo&state=random", githubID))
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func AuthGithub(w http.ResponseWriter, r *http.Request) {
	fmt.Println("redirect")
	url := fmt.Sprintf("https://github.com/login/oauth/access_token?client_id=%s&client_secret=%s&code=%s", utils.GetEnv("GITHUB_ID"), utils.GetEnv("GITHUB_SECRET"), r.FormValue("code"))

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res, _ := json.Marshal("bad request")
		w.Write(res)
	}
	req.Header.Add("Accept", "application/json")

	response, _ := client.Do(req)

	body, _ := ioutil.ReadAll(response.Body)

	requestUser, _ := GetUser(w, r)

	fmt.Println(requestUser.ID)

	accessToken := gjson.GetBytes(body, "access_token")

	models.SetUserToken(strconv.FormatUint(uint64(requestUser.ID), 10), "github_token", accessToken.String())
	CreateWebhook(requestUser.ID ,"JulietteDestang", "test-webhook")
}

func CreateWebhook(userID uint, username string, repository string) {
	if (username == "") || repository == "" {
		return 
	}
	userToken := *models.FindUserToken(userID)
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/hooks", username, repository)

	client := &http.Client{
		Timeout: time.Second * 10,
	}
	fmt.Println(userToken.GithubToken)


	body := []byte(`{"name":"web","active":true,"events":["push","pull_request"],"config":{"url":"https://f206-2a01-cb04-6ff-a300-3b52-f22e-97d5-c899.eu.ngrok.io/webhook/","content_type":"json","insecure_ssl":"0"}}`)


	req, err := http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		fmt.Println("erreur when trying to create webhook")
	}
	req.Header.Add("Authorization", "token " + userToken.GithubToken)
	req.Header.Add("Accept", "application/vnd.github+json")

	response, _ := client.Do(req)
	newbody, _ := ioutil.ReadAll(response.Body)
	fmt.Println(req)
	fmt.Println(response)
	fmt.Println(string(newbody))
}

func Webhook(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
}
