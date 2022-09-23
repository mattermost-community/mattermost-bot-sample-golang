package settings

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
	"time"
)

type Reaction struct {
	Url         string `json:"url"`
	Description string `json:"description"`
}
type (
	Settings struct {
		mu          sync.RWMutex
		settingsUrl string
		settings    CommandSettings
	}

	CommandSettings struct {
		CommandTrigger string              `json:"command_start"`
		Insults        []string            `json:"insults"`
		Quotes         []string            `json:"quotes"`
		Praises        []string            `json:"praises"`
		Reactions      map[string]Reaction `json:"reactions"`
	}
)

func NewSettings(settingsUrl string) (*Settings, error) {
	sc := &Settings{
		settingsUrl: settingsUrl,
	}

	err := sc.LoadSettings()

	return sc, err
}

func (c *Settings) LoadSettings() error {
	var s CommandSettings
	hc := &http.Client{Timeout: 10 * time.Second}

	r, err := hc.Get(c.settingsUrl)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, &s)
	if err != nil {
		return err
	}

	c.mu.Lock()
	c.settings = s

	// this is temporary until we get these pulling from the server
	if len(c.settings.Reactions) == 0 {
		jf, err := os.Open("./reactions.json")
		if err != nil {
			return err
		}
		defer jf.Close()
		b, _ := ioutil.ReadAll(jf)

		json.Unmarshal(b, &c.settings.Reactions)
	}
	c.mu.Unlock()

	return nil
}

func (c *Settings) GetCommandTrigger() string {
	c.mu.RLock()
	commandTrigger := c.settings.CommandTrigger
	c.mu.RUnlock()
	return commandTrigger
}

func (c *Settings) GetInsults() []string {
	c.mu.RLock()
	insults := c.settings.Insults
	c.mu.RUnlock()
	return insults
}

func (c *Settings) GetQuotes() []string {
	c.mu.RLock()
	quotes := c.settings.Quotes
	c.mu.RUnlock()
	return quotes
}

func (c *Settings) GetPraises() []string {
	c.mu.RLock()
	praises := c.settings.Praises
	c.mu.RUnlock()
	return praises
}

func (c *Settings) GetReactions() map[string]Reaction {
	c.mu.RLock()
	reactions := c.settings.Reactions
	c.mu.RUnlock()
	return reactions
}
