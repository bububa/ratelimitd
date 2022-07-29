package app

import (
	"github.com/jinzhu/configor"
	"github.com/urfave/cli/v2"

	"github.com/bububa/ratelimitd/conf"
)

func loadConfigAction(c *cli.Context) error {
	if c.IsSet("debug") {
		config.Debug = c.Bool("debug")
	}
	configPath := c.String("config")
	if err := loadConfig(configPath, &config); err != nil {
		return err
	}
	if c.IsSet("port") {
		config.Port = c.Int("port")
	}
	if c.IsSet("debug") {
		config.Debug = c.Bool("debug")
	}
	if c.IsSet("log") {
		config.LogPath = c.String("log")
	}
	if c.IsSet("prometheus") {
		config.PrometheusEnabled = c.Bool("prometheus")
	}
	return nil
}

func loadConfig(configPath string, cfg *conf.Config) error {
	environment := "production"
	if cfg.Debug {
		environment = "development"
	}
	return configor.New(&configor.Config{
		Verbose:              cfg.Debug,
		ErrorOnUnmatchedKeys: true,
		Environment:          environment,
	}).Load(cfg, configPath)
}

func getServerName(c *cli.Context) error {
	serverName = "ratelimited.xrpc"
	return nil
}
