package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"

	"AREA/pkg/config"
	"AREA/pkg/jobs"
	"AREA/pkg/models"
	"AREA/pkg/utils"
)

// var db * gorm.DB
// var NewUser models.User

var SecretKey = utils.GetEnv("RAPID_API_KEY")

func GetAllUsers(w http.ResponseWriter, r *http.Request){
	newUsers:=models.GetAllUsers()
	res, _ :=json.Marshal(newUsers)
	utils.EnableCors(&w)
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
	utils.EnableCors(&w)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
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
	w.Write(res)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	LoginUser := &models.User{}
	utils.ParseBody(r, LoginUser)

	var user models.User
	user = *models.FindUser(LoginUser.Email)

	if (user.Email == "") {
		fmt.Println("bad email")
		res, _ := json.Marshal("bad email")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(res)
		return
	}
	err := bcrypt.CompareHashAndPassword(user.Password, []byte(LoginUser.Password))
	if (err != nil) {
		fmt.Println("not hash")
		w.WriteHeader(http.StatusBadRequest)
		res, _ := json.Marshal("bad password")
		w.Write(res)
		return
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.ID)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := claims.SignedString([]byte(SecretKey))

	if err != nil {
		fmt.Println("jwt")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cookie := &http.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		Path:     "/",
		HttpOnly: false,
	}

	jobs.AddUserJobsOnLogin(user.ID)
	http.SetCookie(w, cookie)
	res, _ := json.Marshal("sucess")
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

func Logout(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		Path:     "/",
		HttpOnly: false,
	}
	requestUser, _ := GetUser(w, r)
	http.SetCookie(w , cookie)
	fmt.Println(requestUser.ID, requestUser.Firstname)
	jobs.SuprUserJobsOnLogout(requestUser.ID)
	res, _ := json.Marshal("sucess")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetUser(w http.ResponseWriter, r *http.Request) (models.User, error) {
	cookie, cookieErr := r.Cookie("jwt")
	var user models.User
	if (cookieErr != nil) {
		w.WriteHeader(http.StatusBadRequest)
		return user, nil
	}
	token, err := jwt.ParseWithClaims(cookie.Value, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil {
		fmt.Println("ici ?")
		w.WriteHeader(http.StatusBadRequest)
		return user, err
	}
	claims := token.Claims.(*jwt.StandardClaims)
	config.GetDb().Where("id = ?", claims.Issuer).First(&user)
	return user, nil
}


func DoEvery(d time.Duration, f func(time.Time)) {
	for x := range time.Tick(d) {
		f(x)
	}
}

func Helloworld(t time.Time) {
	fmt.Printf("%v: Hello, World!\n", t)
}

func CORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	  w.Header().Add("Access-Control-Allow-Origin", "http://localhost:8081")
	  w.Header().Add("Access-Control-Allow-Credentials", "true")
	  next(w, r)
	}
  }