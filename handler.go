package main

import (
    "regexp"
    "strings"
    "reflect"

    "github.com/mattermost/mattermost-server/v5/model"
    "github.com/pyrousnet/mattermost-golang-bot/commands"
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

func HandleMsgFromChannel (event *model.WebSocketEvent, configuration Configuration) {
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

        commandType := reflect.TypeOf(&commands.Command{})
        commandVal := reflect.ValueOf(&commands.Command{})

        for i := 0; i < commandType.NumMethod(); i++ {
            webSocketClient, err := model.NewWebSocketClient4("wss://" + configuration.Server.HOST + ":" + configuration.Server.PORT, client.AuthToken)
            if err != nil {
                println("We failed to connect to the web socket")
                PrintError(err)
            }
            println("Connected to WS")
            webSocketClient.Listen()

            for resp := range webSocketClient.EventChannel {
                method := commandType.Method(i)
                method.Func.Call([]reflect.Value{commandVal, reflect.ValueOf(resp)})
            }
        }
    }
}
