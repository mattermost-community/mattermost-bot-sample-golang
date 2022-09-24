package commands

import (
	"fmt"
	"math/rand"
	"time"
)

func (h BotCommandHelp) Roll(request BotCommand) (response HelpResponse) {
	return HelpResponse{
		Help:        "Rolls two 6 sided dice for a random response to your query.\n e.g. !roll should I take a break?",
		Description: "Roll some dice!",
	}
}

func (bc BotCommand) Roll(event BotCommand) (response Response, err error) {
	dieSize := 5

	response.Type = "post"

	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	d1 := rand.Intn(dieSize) + 1
	d2 := rand.Intn(dieSize) + 1

	response.Message = fmt.Sprintf("%s rolled a %d and a %d for a total of %d", event.sender, d1, d2, d1+d2)

	return response, nil
}
