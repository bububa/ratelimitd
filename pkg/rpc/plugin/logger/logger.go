package logger

import (
	"github.com/rs/zerolog"
	"github.com/smallnest/rpcx/log"

	"github.com/bububa/ratelimitd/pkg/logger"
)

func init() {
	log.SetLogger(NewLogger(0))
}

func NewLogger(callerSkip int) *Logger {
	return &Logger{
		logger: logger.Logger.With().CallerWithSkipFrameCount(callerSkip).Logger(),
	}
}

type Logger struct {
	logger zerolog.Logger
}

func (l *Logger) Debug(v ...interface{}) {
	l.logger.Debug().Array("v", logger.Array(v)).Msg("rpcx")
}
func (l *Logger) Debugf(format string, v ...interface{}) {
	if len(v) > 0 {
		l.logger.Debug().Msgf(format, v)
		return
	}
	l.logger.Debug().Msg(format)
}

func (l *Logger) Info(v ...interface{}) {
	l.logger.Info().Array("v", logger.Array(v)).Msg("rpcx")
}
func (l *Logger) Infof(format string, v ...interface{}) {
	if len(v) > 0 {
		l.logger.Info().Msgf(format, v)
		return
	}
	l.logger.Info().Msg(format)
}

func (l *Logger) Warn(v ...interface{}) {
	l.logger.Warn().Array("v", logger.Array(v)).Msg("rpcx")
}
func (l *Logger) Warnf(format string, v ...interface{}) {
	if len(v) > 0 {
		l.logger.Warn().Msgf(format, v)
		return
	}
	l.logger.Warn().Msg(format)
}

func (l *Logger) Error(v ...interface{}) {
	l.logger.Error().Array("v", logger.Array(v)).Msg("rpcx")
}
func (l *Logger) Errorf(format string, v ...interface{}) {
	if len(v) > 0 {
		l.logger.Error().Msgf(format, v)
		return
	}
	l.logger.Error().Msg(format)
}
func (l *Logger) Fatal(v ...interface{}) {
	l.logger.Fatal().Array("v", logger.Array(v)).Msg("rpcx")
}
func (l *Logger) Fatalf(format string, v ...interface{}) {
	if len(v) > 0 {
		l.logger.Fatal().Msgf(format, v)
		return
	}
	l.logger.Fatal().Msg(format)
}

func (l *Logger) Panic(v ...interface{}) {
	l.logger.Panic().Array("v", logger.Array(v)).Msg("rpcx")
}
func (l *Logger) Panicf(format string, v ...interface{}) {
	if len(v) > 0 {
		l.logger.Panic().Msgf(format, v)
		return
	}
	l.logger.Panic().Msg(format)
}
