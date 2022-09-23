package commands

import (
	"fmt"
	"strings"
)

func (bc BotCommand) Message(event BotCommand) (response Response, err error) {
	response.Type = "command"

	channelObj, _ := event.mm.GetChannel("town-square")
	response.Channel = channelObj.Id

	target := strings.TrimLeft(event.target, "@")
	response.Message = fmt.Sprintf(`/msg %s %s`, target, event.body)

	return response, nil
}
