package commands

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"

	"github.com/mattermost/mattermost-server/v5/model"
)

const DieSize = 5

func (c Command) HandleRollMsgFromChannel(event *model.WebSocketEvent) {

	// Let's only respond to messaged posted events
	if event.EventType() != model.WEBSOCKET_EVENT_POSTED {
		return
	}

	channelId := event.GetBroadcast().ChannelId
	post := model.PostFromJson(strings.NewReader(event.GetData()["post"].(string)))
	senderName := event.GetData()["sender_name"]

	// Ignore my own events
	if post.UserId == botUser.Id {
		return
	}

	// If message doesn't start with ~roll, ignore it
	if matched, _ := regexp.MatchString(`^~roll(?:$|\W)`, post.Message); matched {
		rand := rand.New(rand.NewSource(time.Now().UnixNano()))
		d1 := rand.Intn(DieSize) + 1
		d2 := rand.Intn(DieSize) + 1

		message := fmt.Sprintf("%s rolled a %d and a %d for a total of %d", senderName, d1, d2, d1+d2)
		SendMsgToChannel(message, channelId, post)
		return
	}
}
