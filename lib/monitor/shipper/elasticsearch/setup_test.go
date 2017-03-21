package elasticsearch_test

import (
	"github.com/ChaosXu/nerv/lib/monitor/model"
	"github.com/ChaosXu/nerv/lib/monitor/shipper/elasticsearch"
	"testing"
)

func TestCreateSchemas(t *testing.T) {
	if metrics, err := model.LoadMetrics("../../../../resources/metrics"); err != nil {
		t.Fatal(err.Error())
	} else {
		elasticsearch.CreateSchemas("localhost:9200", metrics)
	}
}
