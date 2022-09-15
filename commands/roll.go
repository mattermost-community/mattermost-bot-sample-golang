package commands

import (
    "regexp"
    "strings"
    "math/rand"
    "time"
    "fmt"

    "github.com/mattermost/mattermost-server/v5/model"
)

const DieSize = 5

func (c Command) HandleRollMsgFromChannel(event *model.WebSocketEvent) string {
    var post string
    if  p, ok := event.GetData()["post"]; ok {  
        post = model.PostFromJson(strings.NewReader(p.(string))).Message
        //post := model.PostFromJson(strings.NewReader(event.GetData()["post"].(string)))
    } else {
        return ""
    }
    senderName := event.GetData()["sender_name"]
    var message string = ""

    // If message doesn't start with ~roll, ignore it
    if matched, _ := regexp.MatchString(`^!roll(.*)`, post); matched {
        rand := rand.New(rand.NewSource(time.Now().UnixNano()))
        d1 := rand.Intn(DieSize)+1
        d2 := rand.Intn(DieSize)+1

        message = fmt.Sprintf("%s rolled a %d and a %d for a total of %d", senderName, d1, d2, d1+d2)
    }
    return message
}
