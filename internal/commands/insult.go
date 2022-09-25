package commands

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func (bc BotCommand) Insult(event BotCommand) (response Response, err error) {
	insults := event.settings.GetInsults()
	response.Type = "post"
	var index int

	if event.body == "" { //TODO: Check the user.
		response.Type = "dm"
		response.Message = "You must tell me who to insult"

		return response, nil
	}
	arraySize := len(insults)

	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	index = rand.Intn(arraySize)
	response.Message = fmt.Sprintf(`%s`, insults[index])
	response.Message = strings.Replace(response.Message, "{nick}", event.mm.BotUser.Username, -1)
	response.Message = strings.Replace(response.Message, "{0}", event.body, -1)

	return response, nil
}
