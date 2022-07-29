package service

import (
	"github.com/bububa/ratelimitd/conf"
)

func Init(config *conf.Config) {
	configStore = config
}

func Close() {
}
