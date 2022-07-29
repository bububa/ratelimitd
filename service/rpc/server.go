package rpc

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bububa/ratelimitd/pkg/logger"
	"github.com/bububa/ratelimitd/pkg/rpc"
	"github.com/bububa/ratelimitd/service"
	"github.com/smallnest/rpcx/server"
)

func StartServer(ctx context.Context) error {
	config := service.Config()
	logger.Info().Str("service", config.Name).Int("port", config.Port).Str("status", "START").Send()
	limiterService, err := NewService(ctx, config.StoragePath)
	if err != nil {
		logger.Error().Err(err).Send()
		return err
	}
	s := rpc.NewServer(config, limiterService)
	go func() {
		err := s.Serve("tcp", fmt.Sprintf(":%d", config.Port))
		if err != nil && !errors.Is(err, server.ErrServerClosed) {
			logger.Error().Err(err).Send()
		}
	}()
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
	sig := <-ch
	limiterService.Flush(ctx, nil, nil)
	if sig == syscall.SIGHUP {
		logger.Info().Str("service", config.Name).Int("port", config.Port).Str("status", "RESTARTING").Send()
		s.Restart(ctx)
		return nil
	}
	logger.Info().Str("service", config.Name).Int("port", config.Port).Str("status", "STOPPING").Send()
	s.Shutdown(ctx)
	logger.Info().Str("service", config.Name).Int("port", config.Port).Str("status", "EXIT").Send()
	return nil
}
