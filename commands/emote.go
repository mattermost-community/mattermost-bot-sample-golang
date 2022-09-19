package commands

import (
	"fmt"
)

func (bc BotCommand) Emote(event BotCommand) (response Response, err error) {
	response.Type = "command"
	response.Message = fmt.Sprintf(`/me "%s"`, bc.body)

	return response, nil
}
