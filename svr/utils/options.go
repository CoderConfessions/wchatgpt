package utils

import (
	"flag"
	"fmt"
)

func ParseCmd(c *Configuration) error {
	var configFile string
	flag.StringVar(&configFile, "config-file", "", "server config file")
	flag.Parse()

	c.ReadConfig(configFile)
	fmt.Printf("%v\n", c)

	return nil
}
