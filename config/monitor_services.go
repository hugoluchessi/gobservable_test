package config

import (
	m "github.com/hugoluchessi/go-metrics"
	prom "github.com/hugoluchessi/go-metrics/providers/prometheus"
)

type MonitorServices struct {
	*m.MetricService
}

func NewMonitorServices() (*MonitorServices, error) {
	metrics, err := initPrometheusMetrics()

	if err != nil {
		return nil, err
	}

	return &MonitorServices{
		MetricService: metrics,
	}, nil
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
