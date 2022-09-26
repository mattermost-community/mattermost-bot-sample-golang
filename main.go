// Copyright (c) 2016 Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pyrousnet/mattermost-golang-bot/internal/cache"
	"github.com/pyrousnet/mattermost-golang-bot/internal/handler"
	"github.com/pyrousnet/mattermost-golang-bot/internal/mmclient"
	"github.com/pyrousnet/mattermost-golang-bot/internal/settings"
)

func main() {
	//TODO: Set default env to prod
	env := os.Getenv("ENV")
	if env == "" {
		env = "dev"
	}

	cfg, err := settings.GetConfig(env)
	if err != nil {
		log.Fatalln(err.Error())
	}

	mmClient, err := mmclient.NewMMClient(cfg)
	if err != nil {
		log.Fatalln(err.Error())
	}

	botCache := cache.GetCachingMechanism(cfg.Cache.CONN_STR)

	handler, err := handler.NewHandler(mmClient, botCache)
	if err != nil {
		log.Fatalln(err.Error())
	}

	// Lets start listening to some channels via the websocket!
	for {
		ws, err := mmClient.NewWebSocketClient()
		if err != nil {
			log.Fatalf(err.Error())
		}

		fmt.Println("Connected to WS")

		ws.Listen()

		for resp := range ws.EventChannel {
			// We don't want this fella blocking the bot from picking up new events
			go handler.HandleWebSocketResponse(resp)
		}
	}
}
