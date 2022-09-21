package mmclient

import (
	"fmt"
	"os"
	"net/http"
	"time"
	"encoding/json"
	"io/ioutil"

	"github.com/tkanos/gonfig"
)

type Config struct {
	Server struct {
		HOST     string `yaml:"host"`
		PROTOCOL string `yaml:"protocol"`
		PORT     string `yaml:"port"`
	} `yaml:"server"`
	Bot struct {
		SAMPLE_NAME   string `yaml:"sample_name"`
		USER_EMAIL    string `yaml:"user_email"`
		USERNAME      string `yaml:"username"`
		USER_FIRST    string `yaml:"user_first"`
		USER_LAST     string `yaml:"user_last"`
		USER_PASSWORD string `yaml:"user_password"`
		TEAM_NAME     string `yaml:"team_name"`
		LOG_NAME      string `yaml:"log_name"`
		SETTINGS_URL  string `yaml:"settings_url"`
	} `yaml:"bot"`
}

type ChannelList struct {
	id	int
	name	string
}

type Settings struct {
	Server		string		`json: "@server"`
	Password	string		`json: "@password"`
	Port		int		`json: "@port"`
	Secure		int		`json: "@secure"`
	server_validate string		`json: "@server_validate"`
	User		string		`json: "@user"`
	Nick		string		`json: "@nick"`
	Channels	[]ChannelList	`json: "@channels"`
	Admins		[]string	`json: "@admins"`
	Command_start	string		`json: "@command_start"`
	Insults		[]string	`json: "@insults"`
	Quotes		[]string	`json: "@quotes"`
	Praises		[]string	`json: "@praises"`
}

func GetConfig(params ...string) (*Config, error) {
	cfg := Config{}

	//TODO: Set default env to prod
	env := os.Getenv("ENV")
	if env == "" {
		env = "dev"
	}

	// This here is only going to work so long as the bot is started in the main working directory.
	// We might want to set this to pull from some config directory in $GOROOT
	fileName := fmt.Sprintf("./%s_config.yml", env)

	err := gonfig.GetConf(fileName, &cfg)

	return &cfg, err
}

func GetSettings(cfg *Config, target *Settings) (error) {
	httpClient := &http.Client{Timeout: 10 * time.Second}

	r, err := httpClient.Get(cfg.Bot.SETTINGS_URL)

	if err != nil {
		return err
	}
	defer r.Body.Close()

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &target)

	return err
}
