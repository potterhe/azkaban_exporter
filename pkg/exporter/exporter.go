package exporter

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	namespace = "azkaban"
)

type Exporter struct {
	up      *prometheus.Desc
	version *prometheus.Desc
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
		version: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "version"),
			"The version of this azkaban server.",
			[]string{"version"},
			nil,
		),
	}
}

// Describe describes all the metrics exported by the azkaban exporter. It
// implements prometheus.Collector.
func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- e.up
	ch <- e.version
}

// Collect fetches the statistics from the configured azkaban web server, and
// delivers them as Prometheus metrics. It implements prometheus.Collector.
func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	ch <- prometheus.MustNewConstMetric(e.up, prometheus.GaugeValue, 0)
	ch <- prometheus.MustNewConstMetric(e.version, prometheus.GaugeValue, 1, "3.91.0-134-g68e7c718")
}
