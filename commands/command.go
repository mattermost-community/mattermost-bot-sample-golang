package commands
import (
	"github.com/mattermost/mattermost-server/v5/model"
)

type Command struct {
    event *model.WebSocketEvent
}
