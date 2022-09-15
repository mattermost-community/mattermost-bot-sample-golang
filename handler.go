package main

import (
	"reflect"
	"regexp"
	"strings"

	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/pyrousnet/mattermost-golang-bot/commands"
)

type HandlerFunc func(event *model.WebSocketEvent) string

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

	channelId := event.GetBroadcast().ChannelId
	post := model.PostFromJson(strings.NewReader(event.GetData()["post"].(string)))

	// Ignore bot messages
	if post.UserId == botUser.Id {
		return
	}

	if matched, _ := regexp.MatchString(`^!(.*)`, post.Message); matched {
		var messageToSend string = ""

		commandType := reflect.TypeOf(&commands.Command{})
		commandVal := reflect.ValueOf(&commands.Command{})

		for i := 0; i < commandType.NumMethod(); i++ {
			method := commandType.Method(i)
			messageToSend = method.Func.Call([]reflect.Value{commandVal, reflect.ValueOf(event)})[0].Interface().(string)
		}

		if messageToSend != "" {
			SendMsgToChannel(messageToSend, channelId, post)
		}
	}
}
