package jobs

import (
	"fmt"

	"github.com/DisgoOrg/disgohook"

	"AREA/pkg/models"
)

func SendMessage(userID uint, params string) {

	userToken := *models.FindUserToken(userID)

    messageUrl := fmt.Sprintf("%s/%s", userToken.DiscordId, userToken.DiscordToken)

    webhook, _ := disgohook.NewWebhookClientByToken(nil, nil, messageUrl)

    webhook.SendContent(params)
	// Imessage = Imessage
}