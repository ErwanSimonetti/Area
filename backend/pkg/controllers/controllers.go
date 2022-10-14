package controllers

import(
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strconv"
	"AREA/pkg/utils"
	"AREA/pkg/models"
	"github.com/jinzhu/gorm"
	"time"
)

var db * gorm.DB
var NewUser models.User
const SecretKey = "secret"

func GetAllUsers(w http.ResponseWriter, r *http.Request){
	newUsers:=models.GetAllUsers()
	res, _ :=json.Marshal(newUsers)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetUserById(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	userId := vars["userID"]
	ID, err:= strconv.ParseInt(userId,0,0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	userDetails, _:= models.GetUserById(ID)
	res, _ := json.Marshal(userDetails)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


func CreateUser(w http.ResponseWriter, r *http.Request){
	var data map[string]string
	NewUser := &models.User{}
	utils.ParseBody(r, NewUser)
	
	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	NewUser.Password = password
	b := NewUser.CreateUser()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	LoginUser := &models.User{}
	utils.ParseBody(r, LoginUser)

	var user models.User
	user = *models.FindUser(LoginUser.Email)

	if (user.Email == "") {
		res, _ := json.Marshal("bad email")
		w.Write(res)
		return
	}

	if user.Id == 0 {
		w.WriteHeader(http.StatusBadRequest)
	}

	err := bcrypt.CompareHashAndPassword(user.Password, LoginUser.Password)
	if (err != nil) {
		w.WriteHeader(http.StatusBadRequest)
		res, _ := json.Marshal("bad password")
		w.Write(res)
		return
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), //1 day
	})

	token, err := claims.SignedString([]byte(SecretKey))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	cookie := &http.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HttpOnly: true,
	}

	http.SetCookie(w, cookie)
	res, _ := json.Marshal("success")
	w.Write(res)

}

func DeleteUser(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	userId := vars["userId"]
	ID, err := strconv.ParseInt(userId, 0,0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	user := models.DeleteUser(ID)
	res, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateUser(w http.ResponseWriter, r *http.Request){
	var updateUser = &models.User{}
	utils.ParseBody(r, updateUser)
	vars := mux.Vars(r)
	userId := vars["userId"]
	ID, err := strconv.ParseInt(userId, 0,0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	userDetails, db:=models.GetUserById(ID)
	if updateUser.Firstname != ""{
		userDetails.Firstname = updateUser.Firstname
	}
	if updateUser.Lastname != ""{
		userDetails.Lastname = updateUser.Lastname
	}
	if updateUser.Email != ""{
		userDetails.Email = updateUser.Email
	}
	db.Save(&userDetails)
	res, _ := json.Marshal(userDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}