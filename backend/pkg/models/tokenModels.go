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

func SetUserToken(cookie string, column string, token string) {
	fmt.Println(cookie, column, token)
	db.Model(&Token{}).Where("user_id = ?", cookie).Update(column, token)
}