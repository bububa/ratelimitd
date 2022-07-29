package app

import (
	"github.com/urfave/cli/v2"

	"github.com/bububa/ratelimitd/pkg/logger"
	"github.com/bububa/ratelimitd/service/rpc"
)

func beforeCommand(c *cli.Context) error {
	err := getServerName(c)
	if err != nil {
		return err
	}
	sentryWriter, err = InitSentry(&config)
	if err != nil {
		logger.Error().Err(err).Send()
		return err
	}
	InitLogger()
	return nil
}

func RpcStart(c *cli.Context) error {
	rpc.StartServer(c.Context)
	return nil
}
