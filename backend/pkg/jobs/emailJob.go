package jobs

import (
	"fmt"
	"net/smtp"

	"AREA/pkg/models"
	"AREA/pkg/utils"
)

func SendEmail(userID uint, params string) {
	paramsArr := utils.GetParams(params)
	receiver := paramsArr[0]
	message := paramsArr[1]

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