package jobs

import (
	"fmt"

	"github.com/DisgoOrg/disgohook"

	"AREA/pkg/models"
)

func SendMessage(userID uint) {

	userToken := *models.FindUserToken(userID)

    messageUrl := fmt.Sprintf("%s/%s", userToken.DiscordId, userToken.DiscordToken)

    webhook, _ := disgohook.NewWebhookClientByToken(nil, nil, messageUrl)
    msg := "reaction  @everyone"

    webhook.SendContent(msg)
	// Imessage = Imessage
}