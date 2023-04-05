package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	mysqlwrapper "openai-svr/mysql_wrapper"
	"os"
)

type Configuration struct {
	CertFile       string                `json:"cert_file"`
	KeyFile        string                `json:"key_file"`
	OpenaiApiToken string                `json:"openai_api_token"`
	IP             string                `json:"ip"`
	Port           uint                  `json:"port"`
	DB             mysqlwrapper.DBConfig `json:"db"`
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
	if len(c.IP) == 0 {
		return errors.New("IP is not specified")
	}
	if c.Port == 0 {
		return errors.New("Port is not specified")
	}
	if len(c.OpenaiApiToken) == 0 {
		return errors.New("OpenaiApiToken is not specified")
	}
	return nil
}
