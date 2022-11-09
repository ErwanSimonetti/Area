package controllers

import (
	"net/http"
	"strconv"

	"AREA/pkg/models"
)

func AuthEmail(w http.ResponseWriter, r *http.Request) {
  email := r.URL.Query().Get("email")
	password := r.URL.Query().Get("password")
	requestUser, _ := GetUser(w, r)
	models.SetUserToken(strconv.FormatUint(uint64(requestUser.ID), 10), "email", email)
	models.SetUserToken(strconv.FormatUint(uint64(requestUser.ID), 10), "email_password", password)
}