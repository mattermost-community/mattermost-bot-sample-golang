// Copyright (c) 2016 Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package main

import (
    "os"
    "os/signal"
    "regexp"
    "strings"
    "reflect"

    "github.com/mattermost/mattermost-server/v5/model"
    "github.com/pyrousnet/mattermost-golang-bot/commands"
)


var client *model.Client4
var webSocketClient *model.WebSocketClient

var botUser *model.User
var botTeam *model.Team
var debuggingChannel *model.Channel

// Documentation for the Go driver can be found
// at https://godoc.org/github.com/mattermost/platform/model#Client
func main() {
    var connection string
    configuration := GetConfig()
    connection = configuration.Server.PROTOCOL + configuration.Server.HOST

    SetupGracefulShutdown(configuration)

    client = model.NewAPIv4Client(connection)

    // Lets test to see if the mattermost server is up and running
    MakeSureServerIsRunning()

    // lets attempt to login to the Mattermost server as the bot user
    // This will set the token required for all future calls
    // You can get this token with client.AuthToken
    LoginAsTheBotUser(configuration)

    // If the bot user doesn't have the correct information lets update his profile
    UpdateTheBotUserIfNeeded(configuration)

    // Lets find our bot team
    FindBotTeam(configuration)

    // This is an important step.  Lets make sure we use the botTeam
    // for all future web service requests that require a team.
    //client.SetTeamId(botTeam.Id)

    // Lets create a bot channel for logging debug messages into
    CreateBotDebuggingChannelIfNeeded(configuration)
    SendMsgToDebuggingChannel("_"+configuration.Bot.SAMPLE_NAME+" has **started** running_", "")

    RegisterHandlers()

    // Lets start listening to some channels via the websocket!
    for {
        webSocketClient, err := model.NewWebSocketClient4("wss://" + configuration.Server.HOST + ":" + configuration.Server.PORT, client.AuthToken)
        if err != nil {
            println("We failed to connect to the web socket")
            PrintError(err)
        }
        println("Connected to WS")
        webSocketClient.Listen()

        for resp := range webSocketClient.EventChannel {
            HandleWebSocketResponse(resp)
        }
    }
}

// TODO: Is there a way to have each handler register itself??
func RegisterHandlers() {
    //println("Registering roll handler")
    //RegisterHandler(Handler{"roll", HandleRollMsgFromChannel})
    commandType := reflect.TypeOf(&commands.Command{})
    commandVal := reflect.ValueOf(&commands.Command{})

    for i := 0; i < commandType.NumMethod(); i++ {
        method := commandType.Method(i)
        method.Func.Call([]reflect.Value{commandVal, reflect.ValueOf(event)})
    }
}

func MakeSureServerIsRunning() {
    if props, resp := client.GetOldClientConfig(""); resp.Error != nil {
        println("There was a problem pinging the Mattermost server.  Are you sure it's running?")
        PrintError(resp.Error)
        os.Exit(1)
    } else {
        println("Server detected and is running version " + props["Version"])
    }
}

func LoginAsTheBotUser(configuration Configuration) {
    if user, resp := client.Login(
        configuration.Bot.USER_EMAIL, 
        configuration.Bot.USER_PASSWORD); resp.Error != nil {

            println("There was a problem logging into the Mattermost server.  Are you sure ran the setup steps from the README.md?")
            PrintError(resp.Error)
            os.Exit(1)
        } else {
            botUser = user
        }
    }

    func UpdateTheBotUserIfNeeded(configuration Configuration) {
        if botUser.FirstName != configuration.Bot.USER_FIRST || botUser.LastName != configuration.Bot.USER_LAST || botUser.Username != configuration.Bot.USERNAME {
            botUser.FirstName = configuration.Bot.USER_FIRST
            botUser.LastName = configuration.Bot.USER_LAST
            botUser.Username = configuration.Bot.USERNAME

            if user, resp := client.UpdateUser(botUser); resp.Error != nil {
                println("We failed to update the Sample Bot user")
                PrintError(resp.Error)
                os.Exit(1)
            } else {
                botUser = user
                println("Looks like this might be the first run so we've updated the bots account settings")
            }
        }
    }

    func FindBotTeam(configuration Configuration) {
        if team, resp := client.GetTeamByName(configuration.Bot.TEAM_NAME, ""); resp.Error != nil {
            println("We failed to get the initial load")
            println("or we do not appear to be a member of the team '" + configuration.Bot.TEAM_NAME + "'")
            PrintError(resp.Error)
            os.Exit(1)
        } else {
            botTeam = team
        }
    }

    func CreateBotDebuggingChannelIfNeeded(configuration Configuration) {
        println("Attempting to open channel " + configuration.Bot.LOG_NAME)
        if rchannel, resp := client.GetChannelByName(configuration.Bot.LOG_NAME, botTeam.Id, ""); resp.Error != nil {
            println("We failed to get the channels")
            PrintError(resp.Error)
        } else {
            debuggingChannel = rchannel
            return
        }

        // Looks like we need to create the logging channel
        channel := &model.Channel{}
        channel.Name = configuration.Bot.LOG_NAME
        channel.DisplayName = "Debugging For Sample Bot"
        channel.Purpose = "This is used as a test channel for logging bot debug messages"
        channel.Type = model.CHANNEL_OPEN
        channel.TeamId = botTeam.Id
        if rchannel, resp := client.CreateChannel(channel); resp.Error != nil {
            println("We failed to create the channel " + configuration.Bot.LOG_NAME)
            PrintError(resp.Error)
        } else {
            debuggingChannel = rchannel
            println("Looks like this might be the first run so we've created the channel " + configuration.Bot.LOG_NAME)
        }
    }

    func SendMsgToChannel(msg string, channelId string, prePost *model.Post) {
        post := &model.Post{}
        post.ChannelId = channelId
        post.Message = msg
        if prePost.ReplyCount == 0 {
            post.RootId = prePost.Id
        } else {
            post.RootId = prePost.RootId
        }

        if _, resp := client.CreatePost(post); resp.Error != nil {
            println("Failed to send a message to channel " + channelId)
            PrintError(resp.Error)
        }
    }

    func SendMsgToDebuggingChannel(msg string, replyToId string) {
        post := &model.Post{}
        post.ChannelId = debuggingChannel.Id
        post.Message = msg

        post.RootId = replyToId

        if _, resp := client.CreatePost(post); resp.Error != nil {
            println("We failed to send a message to the logging channel")
            PrintError(resp.Error)
        }
    }

    func HandleWebSocketResponse(event *model.WebSocketEvent) {
        HandleMsgFromDebuggingChannel(event)
        HandleMsgFromChannel(event)
    }

    func HandleMsgFromDebuggingChannel(event *model.WebSocketEvent) {
        // If this isn't the debugging channel then lets ingore it
        if event.GetBroadcast().ChannelId != debuggingChannel.Id {
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
            if post.UserId == botUser.Id {
                return
            }

            // if you see any word matching 'alive' then respond
            if matched, _ := regexp.MatchString(`(?:^|\W)alive(?:$|\W)`, post.Message); matched {
                SendMsgToDebuggingChannel("Yes I'm running", post.Id)
                return
            }

            // if you see any word matching 'up' then respond
            if matched, _ := regexp.MatchString(`(?:^|\W)up(?:$|\W)`, post.Message); matched {
                SendMsgToDebuggingChannel("Yes I'm running", post.Id)
                return
            }

            // if you see any word matching 'running' then respond
            if matched, _ := regexp.MatchString(`(?:^|\W)running(?:$|\W)`, post.Message); matched {
                SendMsgToDebuggingChannel("Yes I'm running", post.Id)
                return
            }

            // if you see any word matching 'hello' then respond
            if matched, _ := regexp.MatchString(`(?:^|\W)hello(?:$|\W)`, post.Message); matched {
                SendMsgToDebuggingChannel("Yes I'm running", post.Id)
                return
            }
        }

        SendMsgToDebuggingChannel("I did not understand you!", post.Id)
    }

    func PrintError(err *model.AppError) {
        println("\tError Details:")
        println("\t\t" + err.Message)
        println("\t\t" + err.Id)
        println("\t\t" + err.DetailedError)
    }

    func SetupGracefulShutdown(configuration Configuration) {
        c := make(chan os.Signal, 1)
        signal.Notify(c, os.Interrupt)
        go func() {
            for _ = range c {
                if webSocketClient != nil {
                    webSocketClient.Close()
                }

                SendMsgToDebuggingChannel("_"+configuration.Bot.SAMPLE_NAME+" has **stopped** running_", "")
                os.Exit(0)
            }
        }()
    }
