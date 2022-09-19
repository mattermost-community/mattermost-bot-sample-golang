package commands

import (
	"fmt"
)

func (bc BotCommand) Say(event BotCommand) (response Response, err error) {
	response.Type = "command"
	response.Message = fmt.Sprintf(`/echo "%s" 1`, event.body)

	return response, nil
}
