package main

import (
	"reflect"
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

	channelId := event.GetBroadcast().ChannelId
	post := model.PostFromJson(strings.NewReader(event.GetData()["post"].(string)))

	// Ignore bot messages
	if post.UserId == botUser.Id {
		return
	}

	if matched, _ := regexp.MatchString(`^!(.*)`, post.Message); matched {
		var messageToSend string = ""
		var respType  int = commands.Say

		commandType := reflect.TypeOf(&commands.Command{})
		commandVal := reflect.ValueOf(&commands.Command{})

		for i := 0; i < commandType.NumMethod(); i++ {
			method := commandType.Method(i)
			returns := method.Func.Call([]reflect.Value{commandVal, reflect.ValueOf(event)})
			respType = returns[0].Interface().(int)
			messageToSend = returns[1].Interface().(string)
			if messageToSend != "" {
				break
			}
		}

		println("Received type: %d", respType)
		println("Received message: " + messageToSend)
		if messageToSend != "" {
			if respType == commands.Reply {
				SendMsgToChannel(messageToSend, channelId, post)
			} else if respType == commands.Say {
                // TODO
			} else if respType == commands.Emote {
				SendCmdToChannel(messageToSend, channelId, post)
			}
		}
	}
}
