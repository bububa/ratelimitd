package rpc

import (
	"github.com/smallnest/rpcx/protocol"
	"github.com/smallnest/rpcx/server"
	"github.com/smallnest/rpcx/share"

	"github.com/bububa/ratelimitd/conf"
	"github.com/bububa/ratelimitd/pkg/rpc/codec"
	"github.com/bububa/ratelimitd/pkg/rpc/plugin"
)

func NewServer(cfg *conf.Config, srv interface{}) *server.Server {
	share.Codecs[protocol.SerializeType(4)] = &codec.Msgpack{}
	share.Codecs[protocol.SerializeType(5)] = &codec.Protobuf{}
	s := server.NewServer()
	if !cfg.DisableLogger {
		plugin.AddServerLoggerPlugin(s)
	}
	if cfg.PrometheusEnabled {
		plugin.AddMetricsPlugin(s, cfg)
	}
	s.RegisterName(cfg.Name, srv, "")
	return s
}
