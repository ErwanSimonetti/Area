package models

import (
	"fmt"
	"encoding/json"
	"log"
	
	"gorm.io/datatypes"
	"github.com/jinzhu/gorm"
)

type Token struct {
	gorm.Model
	UserId              uint     `json:"user_id"`
	DiscordId           string   `json:"discord_id"`
	DiscordToken        string   `json:"discord_token"`
	SpotifyToken        string   `json:"spotify_token"`
	SpotifyRefreshToken string   `json:"spotify_refresh_token"`
	Email               string   `json:"email"`
	EmailPassword       string   `json:"email_password"`
	GithubToken         string   `json:"github_token"`
	GithubWebhookIDs    datatypes.JSON `json:"github_webhook_ids"`
}

func (newToken *Token) CreateTokenUser() *Token {
	db.NewRecord(newToken)
	db.Create(&newToken)
	return newToken
}

func FindUserToken(id uint) *Token {
	var getToken Token
	db.Where("user_id = ?", id).Find(&getToken)
	return &getToken
}

func CheckIfConnectedToService(token Token, service string) bool {
	returnValue := false
	switch service {
	case "discord":
		if token.DiscordToken != "" {
			returnValue = true
			break
		}
	case "spotify":
		if token.SpotifyToken != "" {
			returnValue = true
			break
		}
	case "gitHub":
		if token.GithubToken != "" {
			returnValue = true
			break
		}
	}
	return returnValue
}

func SetUserToken(cookie string, column string, token string) {
	fmt.Println(cookie, column, token)
	db.Model(&Token{}).Where("user_id = ?", cookie).Update(column, token)
}

func FindUserByWebhookToken(token string) *Token {
	var getToken Token
	db.Where(datatypes.JSONQuery("github_webhook_ids").Equals(token)).Find(&getToken)
	return &getToken
}

func GetWebhookArray(id uint) datatypes.JSON {
	var webhookArray datatypes.JSON
	db.Select("github_webhook_ids").Where("user_id = ?", id).Find(&webhookArray)
	return webhookArray
}

func UpdateWebhookArray(id uint, newWebhook string) {
	data := GetWebhookArray(id)
	var webhookArray []string
	err := json.Unmarshal([]byte(data), &webhookArray)
	if (err != nil) {
		log.Fatal("can't parse github webhook id")
	}
	webhookArray = append(webhookArray, newWebhook)
	newData, _ := json.Marshal(webhookArray)
	db.Where("user_id = ?", id).Update("github_webhook_ids", newData)
}
