package constants

import (
	"encoding/json"
	"log"
)

type config struct {
	Apps []struct {
		ID  int    `json:"id"`
		URL string `json:"url"`
	} `json:"apps"`
}

var Config *config

func NewConfig(conf []byte) {
	Config = &config{}

	err := json.Unmarshal(conf, Config)

	if err != nil {
		log.Println(err)
	}

	for _, app := range Config.Apps {
		AddAppAgent(app.ID, app.URL)
	}
}
