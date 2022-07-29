package service

import "github.com/bububa/ratelimitd/conf"

var configStore *conf.Config

func Config() *conf.Config {
	return configStore
}
