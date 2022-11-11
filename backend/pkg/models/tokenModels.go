package models

import (
	"fmt"
	// "log"
	
	"github.com/jinzhu/gorm"
)

type GithubWebhook struct {
	gorm.Model
	UserId              uint     `json:"user_id"`
	WebhookID     		string   `json:"webhook_id"`
}

type DiscordWebhook struct {
	gorm.Model
	UserId              uint     `json:"user_id"`
	WebhookID     		string   `json:"webhook_id"`
	WebhookToken     	string   `json:"webhook_token"`
}


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
	// GithubWebhookIDs    GithubWebhook 	 `json:"github_webhook"`
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

func FindUserByDiscordWebhook(id uint) *DiscordWebhook {
	var getToken DiscordWebhook
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

func FindUserByWebhookToken(token string) *GithubWebhook {
	var getToken GithubWebhook
	db.Where("webhook_id = ?", token).Find(&getToken)
	return &getToken
}

func SetGithubWebhook(userId uint, newWebhook string) {
	var newGithubWebhook GithubWebhook
	newGithubWebhook.UserId = userId
	newGithubWebhook.WebhookID = newWebhook
	db.Create(&newGithubWebhook)
	// db.Model(&Token{}).Where("user_id = ?", userId).Update("webhook_id", newWebhook)
}

func SetDiscordWebhook(userId uint, newWebhookID string, newWebhookToken string) {
	var newDiscordWebhook DiscordWebhook
	newDiscordWebhook.UserId = userId
	newDiscordWebhook.WebhookID = newWebhookID
	newDiscordWebhook.WebhookToken = newWebhookToken
	db.Create(&newDiscordWebhook)
	// db.Model(&Token{}).Where("user_id = ?", userId).Update("webhook_id", newWebhook)
}
