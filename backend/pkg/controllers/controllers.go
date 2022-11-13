/** @file controllers.go
 * @brief functions related to users
 *
 * Create, get, delete, login and update users
 */

// @cond

package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"

	"AREA/pkg/models"
	"AREA/pkg/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
)

var db * gorm.DB
// var NewUser models.User
const SecretKey = "secret"

// @endcond


/** Function writing all the users on the response arg (what is it ?)
 *
 * More detailed version
 *
 * @param[in] w http.ResponseWriter
 * @param[in] r *http.Request
 *
 * @return none
 */

func GetAllUsers(w http.ResponseWriter, r *http.Request){
	newUsers:=models.GetAllUsers()
	res, _ :=json.Marshal(newUsers)
	utils.EnableCors(&w)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

/** Get user by his id and does xxx
 *
 * More detailed explanation
 */

func GetUserById(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	userId := vars["userID"]
	ID, err:= strconv.ParseInt(userId,0,0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	userDetails, _:= models.GetUserById(ID)
	res, _ := json.Marshal(userDetails)
	utils.EnableCors(&w)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateUser(w http.ResponseWriter, r *http.Request){
	NewUser := &models.User{}
	utils.ParseBody(r, NewUser)
	password, _ := bcrypt.GenerateFromPassword([]byte(NewUser.Password), 14)

	NewUser.Password = password
	b := NewUser.CreateUser()
	NewUserToken := &models.Token{}
	NewUserToken.UserId = b.ID
	NewUserToken.CreateTokenUser()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	utils.EnableCors(&w)
	w.Write(res)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	utils.EnableCors(&w)
	LoginUser := &models.User{}
	utils.ParseBody(r, LoginUser)

	var user models.User
	user = *models.FindUser(LoginUser.Email)
	userID := *models.FindUserID(LoginUser.Email)
	fmt.Println(userID)
	if (user.Email == "") {
		res, _ := json.Marshal("bad email")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(res)
		return
	}

	err := bcrypt.CompareHashAndPassword(user.Password, []byte(LoginUser.Password))
	if (err != nil) {
		w.WriteHeader(http.StatusBadRequest)
		res, _ := json.Marshal("bad password")
		w.Write(res)
		return
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.ID)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), //1 day
	})

	token, err := claims.SignedString([]byte(SecretKey))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cookie := &http.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HttpOnly: true,
	}

	w.Header().Set("Content-Type","pkglication/json")
	http.SetCookie(w, cookie)
	res, _ := json.Marshal(userID)
	w.WriteHeader(http.StatusOK)
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
	utils.EnableCors(&w)
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
	utils.EnableCors(&w)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}