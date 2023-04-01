package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Configuration struct {
	CertFile       string `json:"cert_file"`
	KeyFile        string `json:"key_file"`
	OpenaiApiToken string `json:"openai_api_token"`
	IP             string `json:"ip"`
	Port           uint   `json:"port"`
}

func NewConfiguration() Configuration {
	return Configuration{
		IP:   "127.0.0.1",
		Port: 8080,
	}
}

func (c *Configuration) ReadConfig(configFile string) error {
	if len(configFile) == 0 {
		return errors.New("configFile is not specified")
	}
	file, _ := os.Open(configFile)
	defer file.Close()
	decoder := json.NewDecoder(file)
	err := decoder.Decode(c)
	if err != nil {
		return errors.New(fmt.Sprintf("read configFile %s failed: %s", configFile, err.Error()))
	}
	return nil
}

func (c *Configuration) ValidateConfig() error {
	return nil
}
