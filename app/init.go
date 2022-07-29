package app

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var (
	GitTag      string // set at compile time with -ldflags
	GitRevision string // set at compile time with -ldflags
	GitSummary  string // set at compile time with -ldflags
)

type DeferFunc func()

func NewApp() (*cli.App, DeferFunc) {
	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Printf("%s-%s@%s\n", serverName, GitRevision, GitTag)
	}

	app := &cli.App{
		Name:  "ratelimited server",
		Usage: "ratelimit service",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "config",
				Aliases:  []string{"c"},
				Usage:    "Load configuration from `FILE`",
				Required: true,
			},
			&cli.BoolFlag{
				Name:  "debug",
				Usage: "Enable debug mode",
				Value: false,
			},
			&cli.StringFlag{
				Name:    "log",
				Aliases: []string{"l"},
				Usage:   "Set Log file to `PATH`",
			},
			&cli.BoolFlag{
				Name:  "prometheus",
				Usage: "Enable prometheus",
			},
		},
		Before: beforeAction,
		After:  afterAction,
		Commands: []*cli.Command{
			{
				Name:     "rpc",
				Usage:    "Start rpc server",
				Category: "Service",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:    "port",
						Aliases: []string{"p"},
						Usage:   "Set server port",
					},
				},
				Before: beforeCommand,
				Action: RpcStart,
			},
		},
	}
	return app, deferFunc
}
