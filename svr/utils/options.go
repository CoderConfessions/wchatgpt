package utils

import (
	"flag"

	"k8s.io/klog/v2"
)

// ParseCmd parse command line for server
func ParseCmd(c *Configuration) error {
	var configFile string
	klog.InitFlags(nil)
	flag.StringVar(&configFile, "config-file", "", "server config file")
	flag.Parse()
	if err := c.ReadConfig(configFile); err != nil {
		return err
	}
	if err := c.ValidateConfig(); err != nil {
		return err
	}

	klog.Info(c)
	return nil
}
