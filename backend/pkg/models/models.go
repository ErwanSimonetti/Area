package models

import (
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

func init() {
	config.Connect()
	db = config.GetDb()
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Token{})
}

func (newUser *User) CreateUser() *User{
	db.NewRecord(newUser)
	db.Create(&newUser)
	return newUser
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
