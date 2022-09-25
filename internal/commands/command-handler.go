package commands

import (
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/pyrousnet/mattermost-golang-bot/internal/mmclient"
	"github.com/pyrousnet/mattermost-golang-bot/internal/settings"

	"github.com/mattermost/mattermost-server/v5/model"
)

type (
	Commands struct {
		availableMethods []Method
		Mm               *mmclient.MMClient
		Settings         *settings.Settings
	}

	Method struct {
		typeOf  reflect.Method
		valueOf reflect.Value
	}

	BotCommand struct {
		body         string
		sender       string
		target       string
		mm           *mmclient.MMClient
		settings     *settings.Settings
		replyChannel *model.Channel
		method       Method
	}

	Response struct {
		Message string
		Type    string
		Channel string
	}
)

func NewCommands(settings *settings.Settings, mm *mmclient.MMClient) *Commands {
	commands := Commands{
		Settings: settings,
		Mm:       mm,
	}

	c := BotCommand{}
	t := reflect.TypeOf(&c)
	v := reflect.ValueOf(&c)
	for i := 0; i < t.NumMethod(); i++ {
		method := Method{
			typeOf:  t.Method(i),
			valueOf: v.Method(i)}

		commands.availableMethods = append(commands.availableMethods, method)
	}

	return &commands
}

func (c *Commands) HandleCommandMsgFromWebSocket(event *model.WebSocketEvent) (Response, error) {
	sender, ok := event.GetData()["sender_name"]
	if ok {
		p, ok := event.GetData()["post"]
		if ok {
			post := model.PostFromJson(strings.NewReader(p.(string))).Message

			bc, err := c.NewBotCommandFromPost(post, sender.(string))
			if err != nil {
				return c.SendErrorResponse(sender.(string), err.Error())
			}

			r, err := c.callCommand(bc)
			if err != nil {
				log.Printf("Error Executing command: %v", err)
				return c.SendErrorResponse(sender.(string), err.Error())
			}

			if r.Channel == "" && bc.replyChannel.Id == "" {
				r.Channel = event.GetBroadcast().ChannelId
			} else {
				r.Channel = bc.replyChannel.Id
			}

			return r, err
		}
		return Response{}, fmt.Errorf("Error: Post not found.")
	}
	return Response{}, fmt.Errorf("Error: Sender not found.")
}

func (c *Commands) SendErrorResponse(target string, message string) (Response, error) {
	replyChannel, _ := c.Mm.GetChannelByName(c.Mm.DebuggingChannel.Name)
	method, _ := c.getMethod("Message")

	bc := BotCommand{
		mm:           c.Mm,
		settings:     c.Settings,
		target:       target,
		replyChannel: replyChannel,
		method:       method,
		body:         message,
	}

	return c.callCommand(bc)
}

func (c *Commands) NewBotCommandFromPost(post string, sender string) (BotCommand, error) {
	ps := strings.Split(post, " ")

	methodName := strings.Title(strings.TrimLeft(ps[0], c.Settings.GetCommandTrigger()))
	ps = append(ps[:0], ps[1:]...)

	method, err := c.getMethod(methodName)
	if err != nil {
		return BotCommand{}, err
	}

	replyChannel := &model.Channel{}
	var rcn string
	if len(ps) > 0 {
		if ps[0] == "in" {
			if len(ps) > 1 {
				rcn = ps[1]
				ps = append(ps[:0], ps[2:]...)

				if rcn != "" {
					c, _ := c.Mm.GetChannel(rcn)
					if c != nil {
						replyChannel = c
					} else {
						log.Default().Println(err)
						return BotCommand{}, fmt.Errorf(`The channel "%s" could not be found.`, rcn)
					}

				}
			}
		}
	}

	body := strings.Join(ps[:], " ")

	return BotCommand{
		mm:           c.Mm,
		settings:     c.Settings,
		body:         body,
		method:       method,
		replyChannel: replyChannel,
	}, nil
}

func (c *Commands) callCommand(botCommand BotCommand) (response Response, err error) {
	f := botCommand.method.valueOf

	in := make([]reflect.Value, 1)
	in[0] = reflect.ValueOf(botCommand)

	var res []reflect.Value
	res = f.Call(in)
	rIface := res[0].Interface()
	if len(res) > 1 {
		e := res[1].Interface()
		if e != nil {
			err = e.(error)
		}
	}

	return rIface.(Response), err
}

func (c *Commands) getMethod(methodName string) (Method, error) {
	for _, m := range c.availableMethods {
		if m.typeOf.Name == methodName {
			return m, nil
		}
	}

	return Method{}, fmt.Errorf("no such command: %s", methodName)
}
