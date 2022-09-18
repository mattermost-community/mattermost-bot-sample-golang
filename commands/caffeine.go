package commands

import (
	"fmt"
	"regexp"
	"strconv"
)

func (c BotCommand) Caffeine(event BotCommand) (response Response, err error) {
	response.Type = "command"

	re := regexp.MustCompile(`^(\d*).*`)
	matched := re.FindStringSubmatch(event.body)

	numShots, err := strconv.Atoi(matched[0])
	if err != nil {
		return response, err
	}

	if numShots != 0 {
		response.Message = fmt.Sprintf("/me walks over to %s and gives them %s shots of caffeine straight into the blood stream.", event.sender, numShots)
	} else {
		response.Message = fmt.Sprintf("/me walks over to %s and gives them a shot of caffeine straight into the blood stream.", event.sender)
	}

	return response, nil
}
