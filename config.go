package main

import (
	"net/url"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type config struct {
	mattermostUserName string
	mattermostTeamName string
	mattermostToken    string
	mattermostChannel  string
	mattermostServer   *url.URL
}

func loadConfig() config {
	var settings config

	settings.mattermostTeamName = os.Getenv("MM_TEAM")
	settings.mattermostUserName = os.Getenv("MM_USERNAME")
	settings.mattermostToken = os.Getenv("MM_TOKEN")
	settings.mattermostChannel = os.Getenv("MM_CHANNEL")
	settings.mattermostServer, _ = url.Parse(os.Getenv("MM_SERVER"))

	return settings
}
