package main
import (
    "fmt"

    "github.com/tkanos/gonfig"
)

type Configuration struct {
    Server struct {
        HOST                string `yaml:"host"`
        PROTOCOL            string `yaml:"protocol"`
        PORT                string `yaml:"port"`
    } `yaml:"server"`
    Bot struct {
        SAMPLE_NAME         string `yaml:"sample_name"`
        USER_EMAIL          string `yaml:"user_email"`
        USERNAME            string `yaml:"username"`
        USER_FIRST          string `yaml:"user_first"`
        USER_LAST           string `yaml:"user_last"`
        USER_PASSWORD       string `yaml:"user_password"`
        TEAM_NAME           string `yaml:"team_name"`
        LOG_NAME            string `yaml:"log_name"`
    } `yaml:"bot"`
}

func GetConfig(params ...string) Configuration {
    configuration := Configuration{}
    env := "dev"
    if len(params) > 0 {
        env = params[0]
    }
    fileName := fmt.Sprintf("./%s_config.yml", env)
    gonfig.GetConf(fileName, &configuration)
    return configuration
}
