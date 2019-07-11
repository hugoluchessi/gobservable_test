package config

import (
	"os"

	log "github.com/hugoluchessi/gobservable/logging"
	m "github.com/hugoluchessi/gobservable/metrics"
	prom "github.com/hugoluchessi/gobservable/metrics/providers/prometheus"
)

type MonitorServices struct {
	*log.ContextLogger
	*m.MetricService
}

func NewMonitorServices() (*MonitorServices, error) {
	metrics, err := initPrometheusMetrics()

	if err != nil {
		return nil, err
	}

	ctxLogger := initLogger()

	return &MonitorServices{
		ContextLogger: ctxLogger,
		MetricService: metrics,
	}, nil
}

func initLogger() *log.ContextLogger {
	zapConfig := log.LoggerConfig{
		Output: os.Stdout,
	}

	logCfgs := []log.LoggerConfig{zapConfig}
	logger := log.NewZapLogger(logCfgs)

	return log.NewContextLogger(logger)
}

func initPrometheusMetrics() (*m.MetricService, error) {
	cfg := m.DefaultConfig("gobservable_test")
	promSink, err := prom.NewSink()

	if err != nil {
		return nil, err
	}

	ms := m.NewMetricService(cfg, promSink)

	if err != nil {
		return nil, err
	}

	return ms, nil
}
