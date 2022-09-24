package commands

import (
	"fmt"
	"strings"
)

func (bc BotCommand) Help(event BotCommand) (response Response, err error) {
	response.Type = "dm"

	channelObj, _ := event.mm.GetChannel(event.mm.DebuggingChannel.Name)
	response.Channel = channelObj.Id

	switch event.body {
	case "react":
		reactions := event.settings.GetReactions()
		event.target = event.sender
		responseMessage := "```\n"

		for i, r := range reactions {
			responseMessage += i + " - " + r.Description + "\n"
		}
		response.Message = fmt.Sprintf("%s```", responseMessage)
	case "roll":
		event.target = event.sender
		target := strings.TrimLeft(event.target, "@")
		responseMessage := "Rolls a single 6 sided die for a random response to your query.\n e.g. !roll should I take a break?"
		response.Message = fmt.Sprintf(`%s %s`, target, responseMessage)
	}

	return response, nil
}
