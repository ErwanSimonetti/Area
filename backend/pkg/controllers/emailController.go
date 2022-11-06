package controllers

import (
  "fmt"
  "net/smtp"
  "AREA/pkg/models"
  "net/http"
  "strconv"
)

func AuthEmail(w http.ResponseWriter, r *http.Request) {

	email := r.URL.Query().Get("email")
	password := r.URL.Query().Get("password")
	requestUser, _ := GetUser(w,r)
	models.SetUserToken(strconv.FormatUint(uint64(requestUser.ID), 10), "email", email)
	models.SetUserToken(strconv.FormatUint(uint64(requestUser.ID), 10), "email_password", password)
}

func SendEmail(userID uint, receiver string, message string) {

  requestUser := *models.FindUserToken(userID)
  from := requestUser.Email
  password := requestUser.EmailPassword

  to := []string{
    receiver,
  }

  smtpHost := "smtp.gmail.com"
  smtpPort := "587"

  messagePayload := []byte(message)

  auth := smtp.PlainAuth("", from, password, smtpHost)
  
  err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, messagePayload)
  if err != nil {
    fmt.Println(err)
    return
  }
}