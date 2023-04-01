package utils

import (
	"flag"

	"k8s.io/klog/v2"
)

func ParseCmd(c *Configuration) error {
	klog.InitFlags(nil)

	var configFile string
	flag.StringVar(&configFile, "config-file", "", "server config file")
	flag.Parse()

	if err := c.ReadConfig(configFile); err != nil {
		return err
	}
	klog.V(4).Infof("%v\n", c)

	if err := c.ValidateConfig(); err != nil {
		return err
	}

	return nil
}
