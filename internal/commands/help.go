package commands

import (
	"fmt"
	"reflect"
	"strings"
)

type (
	BotCommandHelp struct{}

	HelpResponse struct {
		Description string
		Help        string
	}
)

func (bc BotCommand) Help(event BotCommand) (response Response, err error) {
	response.Type = "dm"

	helpMethods := getHelpMethods()
	helpDocs := compileHelpDocs(helpMethods, event)

	bs := strings.Split(event.body, " ")
	if len(bs) > 0 && bs[0] != "" {
		h, ok := helpDocs[strings.Title(bs[0])]
		if ok {
			response.Message = fmt.Sprintf("```\n%s:\n%s\n```", bs[0], h.Help)
		} else {
			response.Message = fmt.Sprintf("Help for '%s' not found.", bs[0])
		}
	} else {
		mess := "```\nAvailable commands:\n"
		for name, helpDoc := range helpDocs {
			mess += strings.ToLower(name) + " - " + helpDoc.Description + "\n"
		}

		response.Message = fmt.Sprintf("%s```", mess)
	}

	return response, nil
}

func getHelpMethods() []Method {
	methods := []Method{}
	c := BotCommandHelp{}
	t := reflect.TypeOf(&c)
	v := reflect.ValueOf(&c)
	for i := 0; i < t.NumMethod(); i++ {
		method := Method{
			typeOf:  t.Method(i),
			valueOf: v.Method(i)}

		methods = append(methods, method)
	}

	return methods
}

func compileHelpDocs(helpMethods []Method, event BotCommand) map[string]HelpResponse {
	response := map[string]HelpResponse{}

	for _, m := range helpMethods {
		f := m.valueOf

		in := make([]reflect.Value, 1)
		in[0] = reflect.ValueOf(event)

		var res []reflect.Value
		res = f.Call(in)
		rIface := res[0].Interface()

		response[m.typeOf.Name] = rIface.(HelpResponse)
	}

	return response
}
