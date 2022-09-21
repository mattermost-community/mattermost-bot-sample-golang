package commands

import (
	"fmt"
	"math/rand"
	"time"
	"strings"
)

func (bc BotCommand) Insult(event BotCommand) (response Response, err error) {
	response.Type = "post"
	var index int

	if event.body == "" { //TODO: Check the user.
		response.Type = "post"
		response.Message = "You must tell me who to insult"
	}
	arraySize := len(event.mm.Settings.Insults)

	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	index = rand.Intn(arraySize)
	response.Message = fmt.Sprintf(`%s`, event.mm.Settings.Insults[index])
	response.Message = strings.Replace(response.Message, "{nick}", event.mm.Settings.Nick, -1)
	response.Message = strings.Replace(response.Message, "{0}", event.body, -1)

	return response, nil
}
