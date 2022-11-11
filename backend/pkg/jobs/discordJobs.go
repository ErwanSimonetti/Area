package jobs

import (
	"fmt"

	"github.com/DisgoOrg/disgohook"

	"AREA/pkg/models"
)

func SendMessage(userID uint, params string) {

	userToken := *models.FindUserByDiscordWebhook(userID)

    messageUrl := fmt.Sprintf("%s/%s", userToken.WebhookID, userToken.WebhookToken)

    webhook, _ := disgohook.NewWebhookClientByToken(nil, nil, messageUrl)

    webhook.SendContent(params)
	// Imessage = Imessage
}