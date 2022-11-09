package jobs

import (
	"fmt"
	"net/smtp"
  
	"AREA/pkg/models"
)

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