package commands

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func (bc BotCommand) Quote(event BotCommand) (response Response, err error) {
	quotes := event.settings.GetQuotes()
	response.Type = "post"
	var index int

	if event.body == "" {
		arraySize := len(quotes)

		rand := rand.New(rand.NewSource(time.Now().UnixNano()))
		index = rand.Intn(arraySize)
	} else {
		index, err = strconv.Atoi(string(event.body[0]))
		if err != nil {
			return response, err
		}
	}
	response.Message = fmt.Sprintf(`%s`, quotes[index])
	response.Message = strings.Replace(response.Message, "{nick}", event.mm.BotUser.Username, -1)
	response.Message = strings.Replace(response.Message, "{0}", event.sender, -1)

	return response, nil
}
