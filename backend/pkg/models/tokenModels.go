package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type Token struct {
	gorm.Model
	UserId uint `json:"userId"`
	DiscordId string `json:"discordId"`
	DiscordToken string `json:"discordToken"`
	SpotifyToken string `json:"spotifyToken"`
	SpotifyRefreshToken string `json:"spotifyRefreshToken"`
	Email string `json:"email"`
	EmailPassword string `json:"emailPassword"`
	GithubToken string `json:"githubToken"`
	WebhookID []string `json:"webhookId"`
}

func (newToken *Token) CreateTokenUser() *Token{
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
		if (token.DiscordToken != "") {
			returnValue = true
			break
		}
	case "spotify":
		if (token.SpotifyToken != "") {
			returnValue = true
			break
		}
	case "gitHub":
		if (token.GithubToken != "") {
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
	db.Where("webhook_id = ?", token).Find(&getToken)
	return &getToken
}

func GetWebhookArray(id uint) []string {
	var webhookArray []string 
	db.Where("user_id = ?", id).Find(&webhookArray)
	return webhookArray
}

func UpdateWebhookArray(id uint , newArray []string) {
	db.Where("user_id = ?", id).Update("webhook_id", newArray)
}