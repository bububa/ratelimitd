package logger

import (
	"context"
	"net"

	"github.com/smallnest/rpcx/protocol"
	"github.com/smallnest/rpcx/server"

	"github.com/bububa/ratelimitd/pkg/logger"
)

type ServerLoggerPlugin struct{}

func (p ServerLoggerPlugin) Register(name string, rcvr interface{}, metadata string) error {
	logger.Info().Str("metadata", metadata).Msg(name)
	return nil
}

func (p ServerLoggerPlugin) Unregister(name string) error {
	logger.Info().Msg("name")
	return nil
}

func (p ServerLoggerPlugin) RegisterFunction(serviceName, fname string, fn interface{}, metadata string) error {
	logger.Info().Str("function", fname).Str("metadata", metadata).Msg(serviceName)
	return nil
}

func (p ServerLoggerPlugin) HandleConnAccept(conn net.Conn) (net.Conn, bool) {
	remoteAddr := conn.RemoteAddr().String()
	logger.Info().Str("remote_addr", remoteAddr).Msg("PostConnAccept")
	return conn, true
}

func (p ServerLoggerPlugin) HandleConnClose(conn net.Conn) (net.Conn, bool) {
	remoteAddr := conn.RemoteAddr().String()
	logger.Info().Str("remote_addr", remoteAddr).Msg("PostConnClose")
	return conn, true
}

func (p ServerLoggerPlugin) PreHandleRequest(ctx context.Context, r *protocol.Message) error {
	clientConn := ctx.Value(server.RemoteConnContextKey).(net.Conn)
	remoteAddr := clientConn.RemoteAddr().String()
	logger.Info().Str("method", r.ServiceMethod).Str("remote_addr", remoteAddr).Msg(r.ServicePath)
	return nil
}

func (p ServerLoggerPlugin) PostWriteResponse(ctx context.Context, req *protocol.Message, res *protocol.Message, err error) error {
	clientConn := ctx.Value(server.RemoteConnContextKey).(net.Conn)
	remoteAddr := clientConn.RemoteAddr().String()
	logger.Info().Str("method", req.ServiceMethod).Str("remote_addr", remoteAddr).Msg(req.ServicePath)
	return nil
}
