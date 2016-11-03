package elasticsearch

import (
	"log"
	"github.com/ChaosXu/nerv/lib/env"
	"github.com/ChaosXu/nerv/lib/monitor/shipper"
)

type ElasticsearchShipper struct {
	server string
	cfg    *env.Properties
}

func NewShipper(cfg *env.Properties) shipper.Shipper {
	address := cfg.GetMapString("shipper", "server", "3334")
	return &ElasticsearchShipper{server:address, cfg:cfg}
}

func (p *ElasticsearchShipper) Init() error {
	return nil
}

func (p *ElasticsearchShipper) Send(v interface{}) {
	//TBD: client pool
	log.Printf("es send: %+v\n", v)
}

