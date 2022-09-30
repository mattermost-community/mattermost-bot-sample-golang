package commands

import (
	"fmt"
)

func (h BotCommandHelp) Say(request BotCommand) (response HelpResponse) {
    response.Help = "Have Bender say something in a specified channel."

    response.Description = "Cause the bot to say something. Usage: '!say in {channel} {text}'"

    return response
}

func (bc BotCommand) Say(event BotCommand) (response Response, err error) {
	response.Type = "command"
	response.Message = fmt.Sprintf(`/echo "%s" 1`, event.body)

	return response, nil
}
