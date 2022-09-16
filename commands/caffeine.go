package commands

import (
	"fmt"
	"regexp"
	"strings"
	"strconv"

	"github.com/mattermost/mattermost-server/v5/model"
)

func (c Command) HandleCaffeineMsgFromChannel(event *model.WebSocketEvent) (string, string) {
	var post string
	var respType string = "command"

	if p, ok := event.GetData()["post"]; ok {
		post = model.PostFromJson(strings.NewReader(p.(string))).Message
		//post := model.PostFromJson(strings.NewReader(event.GetData()["post"].(string)))
	} else {
		return respType, ""
	}
	senderName := event.GetData()["sender_name"]
	var message string = ""

	re := regexp.MustCompile(`^!caffeine\s*(\d*).*`)
	matched := re.FindStringSubmatch(post)
	println("Checking matches")
	fmt.Printf("%+v\n", matched)
	// If message doesn't start with !caffeine, ignore it
	_, err := strconv.Atoi(matched[1])
	if len(matched) > 1 && err == nil {
		message = fmt.Sprintf("/me walks over to %s and gives them %s shots of caffeine straight into the blood stream.", senderName, matched[1])
	} else if matched, _ := regexp.MatchString(`^!caffeine.*`, post); matched {
		message = fmt.Sprintf("/me walks over to %s and gives them a shot of caffeine straight into the blood stream.", senderName)
	}
	return respType, message
}
