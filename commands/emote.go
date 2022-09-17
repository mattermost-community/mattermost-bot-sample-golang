package commands

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/mattermost/mattermost-server/v5/model"
)

func (c Command) HandleEmoteMsgFromChannel(event *model.WebSocketEvent) (int, string) {
	var post string
    var respType int = Emote

	if p, ok := event.GetData()["post"]; ok {
		post = model.PostFromJson(strings.NewReader(p.(string))).Message
		//post := model.PostFromJson(strings.NewReader(event.GetData()["post"].(string)))
	} else {
		return Err, ""
	}
	var message string = ""

	// If message doesn't start with ~roll, ignore it
    re := regexp.MustCompile(`^!emote (.*)`)
    matched := re.FindStringSubmatch(post)
	if len(matched) > 0 {
		message = fmt.Sprintf("/me %s", matched[1])
	}
	return respType, message
}
