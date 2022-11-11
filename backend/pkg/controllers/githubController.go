package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
	"io/ioutil"
	"bytes"
	"encoding/json"
	"strings"

	"github.com/tidwall/gjson"

	"AREA/pkg/utils"
	"AREA/pkg/models"
	"AREA/pkg/jobs"
	
	// "log"
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

	accessToken := gjson.GetBytes(body, "access_token")

	models.SetUserToken(strconv.FormatUint(uint64(requestUser.ID), 10), "github_token", accessToken.String())
	CreateWebhook(requestUser.ID ,"pull_request", "JulietteDestang@@@test-webhook")
	http.Redirect(w, r, "http://localhost:8081/user/services", http.StatusSeeOther)
}

func CreateWebhook(userID uint, action string, params string) {
	split := strings.Split(params, "@@@")
	username := split[0]
	repository := split[1]

	if (username == "") || repository == "" {
		return 
	}
	if (models.CheckExistingGitAction(userID, action)) {
		// log.Fatal("webhook already exist")
		fmt.Println("webhook already exist")
		return
	}

	userToken := *models.FindUserToken(userID)
	
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/hooks", username, repository)

	client := &http.Client{
		Timeout: time.Second * 10,
	}
	body := []byte(`{"name":"web","active":true,"events":["` + action + `"],"config":{"url":"https://f206-2a01-cb04-6ff-a300-3b52-f22e-97d5-c899.eu.ngrok.io/webhook/","content_type":"json","insecure_ssl":"0"}}`)

	req, err := http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		fmt.Println("erreur when trying to create webhook")
	}
	req.Header.Add("Authorization", "token " + userToken.GithubToken)
	req.Header.Add("Accept", "application/vnd.github+json")

	response, _ := client.Do(req)
	newbody, _ := ioutil.ReadAll(response.Body)

	if ((gjson.GetBytes(newbody, "message")).String() == "Bad credentials") {
		fmt.Println("please re log to github")
		return
	}

	if ((gjson.GetBytes(newbody, "message")).String() == "Validation Failed") {
		fmt.Println("webhook already exist")
		return
	}


	webhookID := gjson.GetBytes(newbody, "id")
	models.SetWebhook(userID, webhookID.String())
}

func Webhook(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	webhookID := r.Header.Get("X-Github-Hook-Id")
	webhookEvent := r.Header.Get("X-Github-Event")
	userToken:= *models.FindUserByWebhookToken(webhookID)
	jobs.ExecGithJob(userToken.UserId, webhookEvent)
}
