package commands

import (
	"fmt"
	"github.com/pyrousnet/mattermost-golang-bot/internal/mmclient"
	"log"
	"reflect"
	"strings"

	"github.com/mattermost/mattermost-server/v5/model"
	"golang.org/x/exp/slices"
)

type (
	Commands struct {
		availableMethods []Method
		CommandTrigger   string
		Mm               *mmclient.MMClient
	}

	Method struct {
		typeOf  reflect.Method
		valueOf reflect.Value
	}

	BotCommand struct {
		body   string
		sender string
		mm     *mmclient.MMClient
	}

	Response struct {
		Message string
		Type    string
		Channel string
	}
)

func NewCommands(commandTrigger string, mm *mmclient.MMClient) *Commands {
	commands := Commands{
		CommandTrigger: commandTrigger,
		Mm:             mm,
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

func (c *Commands) HandleCommandMsgFromWebSocket(event *model.WebSocketEvent) Response {
	bc := BotCommand{
		mm: c.Mm,
	}

	if s, ok := event.GetData()["sender_name"]; ok {
		bc.sender = s.(string)
	}

	var post string
	if p, ok := event.GetData()["post"]; ok {
		post = model.PostFromJson(strings.NewReader(p.(string))).Message
	} else {
		log.Println("Error: Post not found.")
		return Response{}
	}

	ps := strings.Split(post, " ")
	methodName := strings.Title(strings.TrimLeft(ps[0], c.CommandTrigger))
	var channel string
	var s string
	if len(ps) > 1 {
		s = fmt.Sprintf("%v", ps[1])
	}
	if len(ps) > 2 {
		channel = fmt.Sprintf("%v", ps[2])
	}

	method, err := c.getMethod(methodName)
	if err != nil {
		return Response{}
	}

	if s == "in" {
		ps = slices.Delete(ps, 0, 3)
		bc.body = strings.Join(ps[0:], " ")
	} else {
		bc.body = strings.Join(ps[1:], " ")
	}

	r, err := c.callCommand(method, bc)
	if s == "in" && channel != "" {
		channelObj, _ := bc.mm.GetChannel(channel)
		if channelObj != nil {
			r.Channel = channelObj.Id
		} else {
			method, err = c.getMethod("Message")
			channelObj, _ = bc.mm.GetChannel(bc.sender)
			bc.body = channel + " is not a channel name."
			r, err = c.callCommand(method, bc)
		}
	}
	if err != nil {
		log.Printf("Error Executing command: %v", err)
	}

	return r
}

func (c *Commands) callCommand(command Method, param BotCommand) (response Response, err error) {
	f := command.valueOf

	in := make([]reflect.Value, 1)
	in[0] = reflect.ValueOf(param)

	var res []reflect.Value
	res = f.Call(in)
	rIface := res[0].Interface()
	return rIface.(Response), nil
}

func (c *Commands) getMethod(methodName string) (Method, error) {
	for _, m := range c.availableMethods {
		if m.typeOf.Name == methodName {
			return m, nil
		}
	}

	return Method{}, fmt.Errorf("no such command: %s", methodName)
}
