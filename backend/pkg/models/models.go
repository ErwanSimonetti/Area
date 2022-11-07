package models

import (
	// "fmt"
	"github.com/jinzhu/gorm"
	"AREA/pkg/config"
)

var db * gorm.DB

type User struct {
	gorm.Model
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
	Email string `json:"email"`
	Password []byte `json:"password"`
}

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

func init() {
	config.Connect()
	db = config.GetDb()
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Token{})
}

func (b *User) CreateUser() *User{
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func (b *Token) CreateTokenUser() *Token{
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func FindUserToken(id uint) *Token {
	var getToken Token
	db.Where("user_id = ?", id).Find(&getToken)
	return &getToken
}

func FindUser(Email string) *User{
	var getUser User
	db.Where("email = ?", Email).Find(&getUser)
	return &getUser
}

func FindUserID(Email string) *uint{
	var getUser User
	db.Where("email = ?", Email).Find(&getUser)
	return &getUser.ID
}

func SetUserToken(cookie string, column string, token string) {
	db.Model(&Token{}).Where("user_id = ?", cookie).Update(column, token)
}

func GetAllUsers() []User{
	var Users []User
	db.Find(&Users)
	return Users
}

func GetUserById(Id int64) (*User, *gorm.DB){
	var getUser User
	db:=db.Where("ID=?", Id).Find(&getUser)
	return &getUser, db
}

func DeleteUser(ID int64) User{
	var User User
	db.Where("ID=?", ID).Delete(User)
	return User
}