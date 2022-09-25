package commands

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func (h BotCommandHelp) Praise(request BotCommand) (response HelpResponse) {
	response.Help = "Give a Bender specific praise from a random list."

	response.Description = "Cause the bot to praise someone. Usage: '!praise {target}'"

	return response
}

func (bc BotCommand) Praise(event BotCommand) (response Response, err error) {
	praises := event.settings.GetPraises()
	response.Type = "post"
	var index int

	if event.body == "" { //TODO: Check the user.
		response.Type = "dm"
		response.Message = "You must tell me who to praise"

		return response, nil
	}
	arraySize := len(praises)

	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	index = rand.Intn(arraySize)
	response.Message = fmt.Sprintf(`%s`, praises[index])
	response.Message = strings.Replace(response.Message, "{nick}", event.mm.BotUser.Username, -1)
	response.Message = strings.Replace(response.Message, "{0}", event.body, -1)

	return response, nil
}
