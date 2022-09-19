package commands

import (
	"fmt"
	"strings"
)

func (c BotCommand) Say(event BotCommand) (response Response, err error) {
	response.Type = "command"

	m := strings.TrimLeft(event.body, " ")
	response.Message = fmt.Sprintf(`/echo "%s" 1`, m)
	return response, nil
}
