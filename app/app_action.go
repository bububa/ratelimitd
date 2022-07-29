package app

import (
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/urfave/cli/v2"

	"github.com/bububa/ratelimitd/service"
)

func beforeAction(c *cli.Context) error {
	if err := loadConfigAction(c); err != nil {
		return err
	}
	if err := InitLogger(); err != nil {
		return err
	}
	service.Init(&config)
	return nil
}

func afterAction(c *cli.Context) error {
	return nil
}

func deferFunc() {
	service.Close()
	if sentryWriter != nil {
		sentryWriter.Close()
		sentry.Recover()
		sentry.Flush(time.Second * 5)
	}
}
