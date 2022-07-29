package app

import (
	"fmt"
	"strings"

	"github.com/getsentry/sentry-go"

	"github.com/bububa/ratelimitd/conf"
	"github.com/bububa/ratelimitd/pkg/logger/zlogsentry"
)

func InitSentry(config *conf.Config) (*zlogsentry.Writer, error) {
	if config.SentryDsn == "" {
		return nil, nil
	}
	environment := "production"
	if strings.Contains(GitSummary, "dirty") {
		environment = "development"
	}
	release := fmt.Sprintf("%s-%s@%s", serverName, GitRevision, GitTag)
	err := sentry.Init(sentry.ClientOptions{
		Dsn:         config.SentryDsn,
		Debug:       config.Debug,
		ServerName:  serverName,
		Release:     release,
		Environment: environment,
		TracesSampler: sentry.TracesSamplerFunc(func(ctx sentry.SamplingContext) sentry.Sampled {
			hub := sentry.GetHubFromContext(ctx.Span.Context())
			if hub == nil {
				return sentry.SampledFalse
			}
			return sentry.UniformTracesSampler(0.7).Sample(ctx)
		}),
	})
	if err != nil {
		return nil, err
	}
	return zlogsentry.New(
		config.SentryDsn,
		zlogsentry.WithServerName(serverName),
		zlogsentry.WithRelease(release),
		zlogsentry.WithEnvironment(environment),
	)
}
