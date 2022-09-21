package commands

import (
	"fmt"
	"strings"
)

func (bc BotCommand) Message(event BotCommand) (response Response, err error) {
	response.Type = "command"
	sender := strings.TrimLeft(event.sender, "@")
	response.Message = fmt.Sprintf(`/msg %s %s`, sender, event.body)

	return response, nil
}
