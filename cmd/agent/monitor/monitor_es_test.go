package monitor_test

import (
	"github.com/ChaosXu/nerv/cmd/agent/monitor"
	"github.com/ChaosXu/nerv/lib/env"
	"testing"
)

func TestElasticsearchShipper(t *testing.T) {
	cfg := env.NewProperties(map[string]interface{}{
		"rpc": map[string]interface{}{
			"port": "4333",
		},
		"shipper": map[string]interface{}{
			"type":   "elasticsearch",
			"server": "localhost:9200",
		},
		"discovery": map[string]interface{}{
			"period": "30",
			"path":   "../../../resources/discovery",
		},
		"monitor": map[string]interface{}{
			"path": "../../../resources/monitor",
		},
		"metrics": map[string]interface{}{
			"path": "../../../resources/metrics",
		},
		"scripts": map[string]interface{}{
			"path": "../../../resources/scripts",
		},
	})

	mock := newMonitorMock()
	mock.expectResource("/host/Linux")
	mock.expectMetric("/host/Linux", "cpu")
	mock.expectMetric("/host/Linux", "memory")

	monitor := monitor.NewMonitor(cfg)
	monitor.Start()

	select {}
}
