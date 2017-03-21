package monitor_test

import (
	"github.com/ChaosXu/nerv/cmd/agent/monitor"
	"github.com/ChaosXu/nerv/lib/env"
	"github.com/ChaosXu/nerv/lib/rpc"
	"log"
	"testing"
	"time"
)

type MonitorMock struct {
	ch        chan map[string]interface{}
	resources map[string]string
	metrics   map[string]string
}

func newMonitorMock() *MonitorMock {
	return &MonitorMock{
		ch:        make(chan map[string]interface{}, 100),
		resources: map[string]string{},
		metrics:   map[string]string{},
	}
}

func (p *MonitorMock) add(v map[string]interface{}) {
	p.ch <- v
}

func (p *MonitorMock) expectResource(resType string) {
	p.resources[resType] = resType
}

func (p *MonitorMock) expectMetric(resType string, metric string) {
	key := resType + "." + metric
	p.metrics[key] = key
}

func (p *MonitorMock) verify(t *testing.T, wait int64) {
	time := time.NewTimer(time.Duration(wait) * time.Second)
	defer time.Stop()

	go func() {
		for c := range p.ch {
			v := c["Type"]
			if v != nil {
				if t, ok := v.(string); ok {
					if t != "" && p.resources[t] != "" {
						p.resources[t] = "ok"
					}
				}
				continue
			}

			v = c["Metric"]
			if v != nil {
				if t, ok := v.(string); ok {
					log.Printf("%s\n", v)
					tags := c["Tags"]
					if ts, ok := tags.(map[string]interface{}); ok {
						rt := ts["resourceType"]
						if r, ok := rt.(string); ok {
							key := r + "." + t
							log.Printf("%s\n", key)
							if t != "" && p.metrics[key] != "" {
								p.metrics[key] = "ok"
							}
						}
					}
				}
				continue
			}
		}
	}()
	<-time.C

	nodatas := map[string]string{}
	count := len(p.resources)
	for k, v := range p.resources {
		if v != "ok" {
			nodatas[k] = v
		} else {
			count--
		}
	}

	if count != 0 {
		t.Fatalf("lose resource:%+v", nodatas)
	} else {
		log.Printf("receive resources: %+v", p.resources)
	}

	nodatas = map[string]string{}
	count = len(p.metrics)
	for k, v := range p.metrics {
		if v != "ok" {
			nodatas[k] = v
		} else {
			count--
		}
	}

	if count != 0 {
		t.Fatalf("lose metrics:%+v", nodatas)
	} else {
		log.Printf("receive metrics: %+v", p.metrics)
	}
}

//MonitorPublisher for test
type MonitorPublisher struct {
	mock *MonitorMock
}

func (p *MonitorPublisher) Publish(v map[string]interface{}, reply *string) error {
	log.Printf("test publish: %+v\n", v)
	p.mock.add(v)
	return nil
}

func TestStartStop(t *testing.T) {
	cfg := env.NewProperties(map[string]interface{}{
		"rpc": map[string]interface{}{
			"port": "4333",
		},
		"shipper": map[string]interface{}{
			"type":   "rpc",
			"server": "localhost:4333",
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
	mock.expectResource("/db/Mysql")
	mock.expectResource("/middleware/Nginx")
	mock.expectResource("/middleware/Tomcat")
	mock.expectResource("/nerv/Agent")

	mock.expectMetric("/host/Linux", "cpu")
	mock.expectMetric("/host/Linux", "memory")

	go startServer(t, cfg, mock)

	monitor := monitor.NewMonitor(cfg)
	monitor.Start()

	mock.verify(t, 10)
}

func startServer(t *testing.T, cfg *env.Properties, mock *MonitorMock) {
	server := &MonitorPublisher{mock}
	rpc.Register(server)
	if err := rpc.Start(cfg); err != nil {
		t.Fatal(err.Error())
	}
}
