package exporter

import (
	"fmt"

	"github.com/potterhe/azkaban_exporter/pkg/azkaban"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	namespace = "azkaban"
)

type Exporter struct {
	client *azkaban.Client

	up      *prometheus.Desc
	version *prometheus.Desc
}

// New returns an initialized exporter.
func New(server string) *Exporter {
	return &Exporter{
		client: &azkaban.Client{
			Server: server,
		},
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

	// todo 连接服务器
	ch <- prometheus.MustNewConstMetric(e.up, prometheus.GaugeValue, 0)

	resp, err := e.client.Status()
	fmt.Println(resp, err)
	if err == nil {
		ch <- prometheus.MustNewConstMetric(e.version, prometheus.GaugeValue, 1, resp.Version)
	}

}
