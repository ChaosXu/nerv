package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

// Proxy config
type Proxy struct {
	rp *httputil.ReverseProxy
}

func NewProxy(remoteUrl string) (*Proxy, error) {
	url, err := url.Parse(remoteUrl)
	if err != nil {
		return nil, err
	}
	return &Proxy{rp: httputil.NewSingleHostReverseProxy(url)}, nil
}

func (p *Proxy) Handle(w http.ResponseWriter, r *http.Request) {
	p.rp.ServeHTTP(w, r)
}
