package commands

import "fmt"

func (bc BotCommand) React(event BotCommand) (response Response, err error) {
	reactions := event.settings.GetReactions()
	if r, ok := reactions[event.body]; ok {
		response.Type = "command"
		response.Message = fmt.Sprintf(`/echo "%s" 1`, r.Url)
	} else {
		response.Type = "post"
		err = fmt.Errorf("Response key '%s' not found.", event.body)
		response.Message = fmt.Sprintf("%s", err)
	}

	return response, err
}
