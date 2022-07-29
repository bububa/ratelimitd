package app

import (
	"github.com/bububa/ratelimitd/conf"
	"github.com/bububa/ratelimitd/pkg/logger/zlogsentry"
)

var (
	// config global configuration
	config conf.Config
	// serverName service name
	serverName string
	// sentryWriter
	sentryWriter *zlogsentry.Writer
)
