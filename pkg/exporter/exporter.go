package exporter

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	namespace = "azkaban"
)

type Exporter struct {
	up *prometheus.Desc
}

// New returns an initialized exporter.
func New() *Exporter {
	return &Exporter{
		up: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "up"),
			"Could the azkaban web server be reached.",
			nil,
			nil,
		),
	}
}

// Describe describes all the metrics exported by the azkaban exporter. It
// implements prometheus.Collector.
func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- e.up
}

// Collect fetches the statistics from the configured azkaban web server, and
// delivers them as Prometheus metrics. It implements prometheus.Collector.
func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	ch <- prometheus.MustNewConstMetric(e.up, prometheus.GaugeValue, 0)
}
