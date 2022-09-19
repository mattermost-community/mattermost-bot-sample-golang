package commands

import (
	"fmt"
)

func (bc BotCommand) Emote(event BotCommand) (response Response, err error) {
	response.Type = "command"
	response.Message = fmt.Sprintf(`/me %s`, event.body)

	return response, nil
}
