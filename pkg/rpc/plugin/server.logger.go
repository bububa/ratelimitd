package plugin

import (
	"github.com/smallnest/rpcx/server"

	"github.com/bububa/ratelimitd/pkg/rpc/plugin/logger"
)

func AddServerLoggerPlugin(s *server.Server) error {
	r := &logger.ServerLoggerPlugin{}
	s.Plugins.Add(r)
	return nil
}
