package commands

import (
	"fmt"
)

func (bc BotCommand) Say(event BotCommand) (response Response, err error) {
    response.Type = "command"
    fmt.Println("body: "+bc.body)
    response.Message = fmt.Sprintf(`/echo "%s" 1`, bc.body)
	return response, nil
}
