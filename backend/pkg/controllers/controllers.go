package controllers

import(
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"AREA/pkg/utils"
	"AREA/pkg/models"
)

var NewUser models.User

func GetUser(w http.ResponseWriter, r *http.Request){
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
	CreateUser := &models.User{}
	utils.ParseBody(r, CreateUser)
	b := CreateUser.CreateUser()
	// fmt.Println(b)
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// func CreateBook(w http.ResponseWriter, r *http.Request){
// 	CreateBook := &models.Book{}
// 	utils.ParseBody(r, CreateBook)
// 	b:= CreateBook.CreateBook()
// 	res, _ := json.Marshal(b)
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)
// }

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