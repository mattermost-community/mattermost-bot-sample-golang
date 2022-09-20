package commands

import (
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/mattermost/mattermost-server/v5/model"
)

type (
	Commands struct {
		availableMethods []Method
		CommandTrigger   string
	}

	Method struct {
		typeOf  reflect.Method
		valueOf reflect.Value
	}

	BotCommand struct {
		body    string
		sender  string
		channel string
	}

	Response struct {
		Message string
		Type    string
	}
)

func NewCommands(commandTrigger string) *Commands {
	commands := Commands{
		CommandTrigger: commandTrigger,
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
	bc := BotCommand{}

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
	s := fmt.Sprintf("%c", ps[1])
	channel := fmt.Sprintf("%c", ps[2])
	if s == "in" {
		channelId, _ := main.GetChannel(channel)
	} else {
		bc.body = strings.Join(ps[1:], " ")
	}

	method, err := c.getMethod(methodName)
	if err != nil {
		return Response{}
	}

	r, err := c.callCommand(method, bc)
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
