package models

import (
	"github.com/jinzhu/gorm"
	"AREA/pkg/config"
)

var db * gorm.DB

type User struct {
	gorm.Model
	Firstname string `gorm:""json:"firstname"`
	Lastname string `json:"lastname"`
	Email string `json:"email"`
}

func init() {
	config.Connect()
	db = config.GetDb()
	db.AutoMigrate(&User{})
}

func (b *User) CreateUser() *User{
	db.NewRecord(b)
	db.Create(&b)
	return b
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