package commands

import (
	"fmt"
)

func (bc BotCommand) Apb(event BotCommand) (response Response, err error) {
	response.Type = "command"
    // TODO - Check if the user exists.
	response.Message = fmt.Sprintf(`/me sends out the blood hounds to find %s`, event.body)

	return response, nil
}
