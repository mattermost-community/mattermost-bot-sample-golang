package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/pyrousnet/mattermost-golang-bot/commands"
)

type HandlerFunc func(event *model.WebSocketEvent) (string, string)

type Handler struct {
	command string
	method  HandlerFunc
}

var HandlerList []Handler

func RegisterHandler(handler Handler) {
	HandlerList = append(HandlerList, handler)
}

func HandleMsgFromChannel(event *model.WebSocketEvent, configuration Configuration) {
	//Only handle messaged posted events
	if event.EventType() != model.WEBSOCKET_EVENT_POSTED {
		return
	}

	// TODO: Move this to settings
	commandTrigger := "!"

	cmds := commands.NewCommands(commandTrigger)

	channelId := event.GetBroadcast().ChannelId
	post := model.PostFromJson(strings.NewReader(event.GetData()["post"].(string)))

	// Ignore bot messages
	if post.UserId == botUser.Id {
		return
	}

	pattern := fmt.Sprintf(`^%s(.*)`, commandTrigger)

	if ok, _ := regexp.MatchString(pattern, post.Message); ok {
		response := cmds.HandleCommandMsgFromWebSocket(event)

		if response.Message != "" {
			switch response.Type {
			case "post":
				SendMsgToChannel(response.Message, channelId, post)
			case "command":
				SendCmdToChannel(response.Message, channelId, post)
			}
		}
	}
}
