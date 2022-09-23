package commands

import (
	"fmt"
	"strings"
)

func (bc BotCommand) Help(event BotCommand) (response Response, err error) {
	response.Type = "command"

	channelObj, _ := event.mm.GetChannel("town-square")
	response.Channel = channelObj.Id

	switch event.body {
	case "react":
		reactions := event.settings.GetReactions()
		event.target = event.sender
		responseMessage := "```\n"

		for i, r := range reactions {
			responseMessage += i + " - " + r.Description + "\n"
		}

		target := strings.TrimLeft(event.target, "@")
		response.Message = fmt.Sprintf(`/msg %s %s`, target, responseMessage)
	}

	return response, nil
}
