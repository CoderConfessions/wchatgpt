package utils

import (
	"flag"
	"fmt"
	"path/filepath"
	"time"

	"k8s.io/klog/v2"
)

// ParseCmd parse command line for server
func ParseCmd(c *Configuration) error {
	var configFile string
	flag.StringVar(&configFile, "config-file", "", "server config file")
	flag.Parse()
	if err := c.ReadConfig(configFile); err != nil {
		return err
	}
	if err := c.ValidateConfig(); err != nil {
		return err
	}
	initLogFlag(c)

	klog.Info(c)
	klog.V(4).Info(c)
	return nil
}

// initLogFlag init klog flags
func initLogFlag(c *Configuration) {
	klog.InitFlags(nil)
	flag.Set("log_file",
		filepath.Join(c.Log.Path,
			fmt.Sprintf("wchatgpt-%s.log", time.Now().Format("20060102-150405"))))
	flag.Set("alsologtostderr", c.Log.Alsologtostderr)
	flag.Set("logtostderr", c.Log.Logtostderr)
	flag.Set("v", c.Log.V) // why not take affect?
}
