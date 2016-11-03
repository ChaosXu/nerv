package shipper

import (
	"log"
	"github.com/ChaosXu/nerv/lib/env"
)

type ElasticsearchShipper struct {
	server string
	cfg    *env.Properties
}

func NewElasticsearchShipper(cfg *env.Properties) Shipper {
	address := cfg.GetMapString("shipper", "server", "3334")
	return &ElasticsearchShipper{server:address, cfg:cfg}
}

func (p *ElasticsearchShipper) Send(v interface{}) {
	//TBD: client pool
	log.Printf("es send: %+v\n", v)
}
