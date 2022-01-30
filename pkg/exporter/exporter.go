package exporter

import (
	"fmt"
	"strconv"

	"github.com/potterhe/azkaban_exporter/pkg/azkaban"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	namespace = "azkaban"
)

type Exporter struct {
	client *azkaban.Client

	up              *prometheus.Desc
	version         *prometheus.Desc
	databaseUp      *prometheus.Desc
	executorStatus  *prometheus.Desc
	usedMemoryBytes *prometheus.Desc
	xmxBytes        *prometheus.Desc
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
		databaseUp: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "database_up"),
			"Could the database be reached.",
			nil,
			nil,
		),
		executorStatus: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "executor_status"),
			"Executor status",
			[]string{"host", "is_active"},
			nil,
		),
		usedMemoryBytes: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "used_memory_bytes"),
			"Azkaban web server used memory bytes",
			nil,
			nil,
		),
		xmxBytes: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, "", "xmx_bytes"),
			"Azkaban web server xmx bytes",
			nil,
			nil,
		),
	}
}

// Describe describes all the metrics exported by the azkaban exporter. It
// implements prometheus.Collector.
func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- e.up
	ch <- e.version
	ch <- e.databaseUp
	ch <- e.executorStatus
	ch <- e.usedMemoryBytes
}

// Collect fetches the statistics from the configured azkaban web server, and
// delivers them as Prometheus metrics. It implements prometheus.Collector.
func (e *Exporter) Collect(ch chan<- prometheus.Metric) {

	// todo 连接服务器
	ch <- prometheus.MustNewConstMetric(e.up, prometheus.GaugeValue, 0)

	status, err := e.client.Status()
	fmt.Println(status, err)
	if err == nil {
		ch <- prometheus.MustNewConstMetric(e.version, prometheus.GaugeValue, 1, status.Version)
		var iDatabaseUp float64 = 0
		if status.IsDatabaseUp {
			iDatabaseUp = 1
		}
		ch <- prometheus.MustNewConstMetric(e.databaseUp, prometheus.GaugeValue, iDatabaseUp)

		for _, executor := range status.ExecutorStatusMap {
			ch <- prometheus.MustNewConstMetric(e.executorStatus, prometheus.GaugeValue, 1, executor.Host, strconv.FormatBool(executor.IsActive))
		}
		ch <- prometheus.MustNewConstMetric(e.usedMemoryBytes, prometheus.GaugeValue, float64(status.UsedMemory))
		ch <- prometheus.MustNewConstMetric(e.xmxBytes, prometheus.GaugeValue, float64(status.Xmx))
	}

}
