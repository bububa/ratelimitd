package plugin

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
	"github.com/smallnest/rpcx/protocol"
	"github.com/smallnest/rpcx/server"

	"github.com/bububa/ratelimitd/conf"
	"github.com/bububa/ratelimitd/pkg/logger"
)

const metricsNamespace = "ratelimitd"

var (
	labels = []string{"method"}

	uptime = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: metricsNamespace,
			Name:      "rpcx_uptime",
			Help:      "RPCX service uptime.",
		}, nil,
	)

	reqCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: metricsNamespace,
			Name:      "rpcx_request_count_total",
			Help:      "Total number of RPCX requests made.",
		}, labels,
	)

	reqDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: metricsNamespace,
			Name:      "rpcx_request_duration_seconds",
			Help:      "RPCX request latencies in seconds.",
		}, labels,
	)

	reqSizeBytes = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Namespace: metricsNamespace,
			Name:      "rpcx_request_size_bytes",
			Help:      "RPCX request sizes in bytes.",
		}, labels,
	)

	respSizeBytes = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Namespace: metricsNamespace,
			Name:      "rpcx_response_size_bytes",
			Help:      "HTTP request sizes in bytes.",
		}, labels,
	)
)

func AddMetricsPlugin(s *server.Server, config *conf.Config) error {
	if config.PrometheusGateway == "" {
		return errors.New("invalid config")
	}
	registry := prometheus.NewRegistry()
	registry.MustRegister(uptime, reqCount, reqDuration, reqSizeBytes, respSizeBytes)
	registry.MustRegister(prometheus.NewGoCollector(), prometheus.NewProcessCollector(prometheus.ProcessCollectorOpts{ReportErrors: true}))
	job := fmt.Sprintf("ratelimitd.%s", config.Name)
	pusher := push.New(config.PrometheusGateway, job).Gatherer(registry)
	go recordUptime()
	go pushMetrics(pusher)
	r := NewServerPromPlugin()
	s.Plugins.Add(r)
	return nil
}

type promReporter struct {
	method    string
	startTime time.Time
}

func newPromReporter(method string) *promReporter {
	return &promReporter{
		method:    method,
		startTime: time.Now(),
	}
}

func (r *promReporter) ReceivedMessage(reqSize float64) {
	reqCount.WithLabelValues(r.method).Inc()
	reqSizeBytes.WithLabelValues(r.method).Observe(reqSize)
}

func (r *promReporter) Handled(respSize float64) {
	reqDuration.WithLabelValues(r.method).Observe(time.Since(r.startTime).Seconds())
	respSizeBytes.WithLabelValues(r.method).Observe(respSize)
}

type ServerPromPlugin struct {
	mp map[context.Context]*promReporter
	sync.RWMutex
}

func NewServerPromPlugin() *ServerPromPlugin {
	return &ServerPromPlugin{
		mp: make(map[context.Context]*promReporter),
	}
}

func (p *ServerPromPlugin) PreHandleRequest(ctx context.Context, r *protocol.Message) error {
	monitor := newPromReporter(r.ServiceMethod)
	var reqSize float64
	if r != nil {
		reqSize = float64(len(r.Payload))
	}
	monitor.ReceivedMessage(reqSize)

	p.Lock()
	defer p.Unlock()
	p.mp[ctx] = monitor
	return nil
}

func (p *ServerPromPlugin) PostWriteResponse(ctx context.Context, req *protocol.Message, res *protocol.Message, err error) error {
	var monitor *promReporter
	p.RLock()
	monitor = p.mp[ctx]
	p.RUnlock()

	if monitor == nil {
		return nil
	}
	var respSize float64
	if res != nil {
		respSize = float64(len(res.Payload))
	}

	monitor.Handled(respSize)

	p.Lock()
	defer p.Unlock()
	delete(p.mp, ctx)
	return nil
}

func recordUptime() {
	for range time.Tick(time.Second) {
		uptime.WithLabelValues().Inc()
	}
}

func pushMetrics(pusher *push.Pusher) {
	for range time.Tick(5 * time.Second) {
		err := pusher.Push()
		if err != nil {
			logger.Error().Err(err).Send()
		}
	}
}
