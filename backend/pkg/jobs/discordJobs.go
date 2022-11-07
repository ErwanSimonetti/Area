package jobs

import (
	"fmt"

	"github.com/DisgoOrg/disgohook"

	"AREA/pkg/models"
)

func SendMessage(userID uint) {

	userToken := *models.FindUserToken(userID)
	fmt.Println(userToken.DiscordId)
	fmt.Println(userToken.DiscordToken)

    messageUrl := fmt.Sprintf("%s/%s", userToken.DiscordId, userToken.DiscordToken)

    webhook, _ := disgohook.NewWebhookClientByToken(nil, nil, messageUrl)
    msg := "reaction  @everyone"

    Imessage, _ := webhook.SendContent(msg)
	Imessage = Imessage
	fmt.Println("reaction @everyone")
}