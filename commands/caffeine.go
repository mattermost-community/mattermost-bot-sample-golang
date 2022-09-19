package commands

import (
	"fmt"
	"strconv"
)

func (c BotCommand) Caffeine(event BotCommand) (response Response, err error) {
	response.Type = "command"

	if event.body != "" {
		numShots, err := strconv.Atoi(string(event.body[0]))
		if err == nil {
			fmt.Printf("%+v\n", numShots)
			response.Message = fmt.Sprintf("/me walks over to %s and gives them %d shots of caffeine straight into the blood stream.", event.sender, numShots)
		}
	} else {
		response.Message = fmt.Sprintf("/me walks over to %s and gives them a shot of caffeine straight into the blood stream.", event.sender)
	}

	return response, nil
}
