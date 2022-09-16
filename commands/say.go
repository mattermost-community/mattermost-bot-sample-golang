package commands

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/mattermost/mattermost-server/v5/model"
)

func (c Command) HandleSayMsgFromChannel(event *model.WebSocketEvent) string {
	var post string
	if p, ok := event.GetData()["post"]; ok {
		post = model.PostFromJson(strings.NewReader(p.(string))).Message
		//post := model.PostFromJson(strings.NewReader(event.GetData()["post"].(string)))
	} else {
		return ""
	}
	var message string = ""

	// If message doesn't start with ~roll, ignore it
    re := regexp.MustCompile(`^!say (.*)`)
    matched := re.FindStringSubmatch(post)
	if len(matched) > 0 {
		message = fmt.Sprintf("%s", matched[1])
	}
	return message
}
