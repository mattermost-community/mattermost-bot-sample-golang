package main

import (
	"regexp"
	"strings"

	"github.com/mattermost/mattermost-server/v5/model"
)

type HandlerFunc func(event *model.WebSocketEvent) string

type Handler struct {
	command	string
	method	HandlerFunc
}

var HandlerList []Handler

func RegisterHandler(handler Handler) {
	HandlerList = append(HandlerList, handler)
}

func HandleMsgFromChannel (event *model.WebSocketEvent) {
	//Only handle messaged posted events
	if event.EventType() != model.WEBSOCKET_EVENT_POSTED {
		return
	}

	channelId := event.GetBroadcast().ChannelId
	post := model.PostFromJson(strings.NewReader(event.GetData()["post"].(string)))

	// Ignore bot messages
	if post.UserId == botUser.Id {
		return
	}

	if matched, _ := regexp.MatchString(`^!(.*)`, post.Message); matched {
		var handlerResponse string = ""

		for _, handler := range HandlerList {
			handlerResponse = handler.method(event);
		 	if handlerResponse != "" {
				 break
		 	}
		}

		if handlerResponse != "" {
			SendMsgToChannel(handlerResponse, channelId, post)
		}
	}
}
