package handler

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/pyrousnet/mattermost-golang-bot/internal/commands"
	"github.com/pyrousnet/mattermost-golang-bot/internal/mmclient"

	"github.com/mattermost/mattermost-server/v5/model"
)

type Handler struct {
	mm *mmclient.MMClient
}

func NewHandler(mm *mmclient.MMClient) *Handler {
	return &Handler{
		mm: mm,
	}
}

func (h *Handler) HandleWebSocketResponse(event *model.WebSocketEvent) {
	h.HandleMsgFromDebuggingChannel(event)
	h.HandleMsgFromChannel(event)
}

func (h *Handler) HandleMsgFromChannel(event *model.WebSocketEvent) {
	//Only handle messaged posted events
	if event.EventType() != model.WEBSOCKET_EVENT_POSTED {
		return
	}

	// TODO: Move this to settings
	commandTrigger := "!"

	cmds := commands.NewCommands(commandTrigger, h.mm)

	channelId := event.GetBroadcast().ChannelId
	post := model.PostFromJson(strings.NewReader(event.GetData()["post"].(string)))

	// Ignore bot messages
	if post.UserId == h.mm.BotUser.Id {
		return
	}

	pattern := fmt.Sprintf(`^%s(.*)`, commandTrigger)

	if ok, _ := regexp.MatchString(pattern, post.Message); ok {
		response := cmds.HandleCommandMsgFromWebSocket(event)
		println("here " + response.Channel)
		if "" == response.Channel {
			response.Channel = channelId
		}

		if response.Message != "" {
			switch response.Type {
			case "post":
				h.mm.SendMsgToChannel(response.Message, response.Channel, post)
			case "command":
				h.mm.SendCmdToChannel(response.Message, response.Channel, post)
			}
		}
	}
}

func (h *Handler) HandleMsgFromDebuggingChannel(event *model.WebSocketEvent) {
	// If this isn't the debugging channel then lets ingore it
	if event.GetBroadcast().ChannelId != h.mm.DebuggingChannel.Id {
		return
	}

	// Lets only reponded to messaged posted events
	if event.EventType() != model.WEBSOCKET_EVENT_POSTED {
		return
	}

	println("responding to debugging channel msg")

	post := model.PostFromJson(strings.NewReader(event.GetData()["post"].(string)))
	if post != nil {

		// ignore my events
		if post.UserId == h.mm.BotUser.Id {
			return
		}

		// if you see any word matching 'alive' then respond
		if matched, _ := regexp.MatchString(`(?:^|\W)alive(?:$|\W)`, post.Message); matched {
			h.mm.SendMsgToDebuggingChannel("Yes I'm running", post.Id)
			return
		}

		// if you see any word matching 'up' then respond
		if matched, _ := regexp.MatchString(`(?:^|\W)up(?:$|\W)`, post.Message); matched {
			h.mm.SendMsgToDebuggingChannel("Yes I'm running", post.Id)
			return
		}

		// if you see any word matching 'running' then respond
		if matched, _ := regexp.MatchString(`(?:^|\W)running(?:$|\W)`, post.Message); matched {
			h.mm.SendMsgToDebuggingChannel("Yes I'm running", post.Id)
			return
		}

		// if you see any word matching 'hello' then respond
		if matched, _ := regexp.MatchString(`(?:^|\W)hello(?:$|\W)`, post.Message); matched {
			h.mm.SendMsgToDebuggingChannel("Yes I'm running", post.Id)
			return
		}
	}

	h.mm.SendMsgToDebuggingChannel("I did not understand you!", post.Id)
}
