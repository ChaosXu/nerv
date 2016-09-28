package middleware

import (
	"net/http"
	"golang.org/x/net/context"
	"github.com/pressly/chi"
	"strings"
	"net/url"
)

const (
	RequestParams = "params"
)

// ParamsParser parse path variables and query variables from url
func ParamsParser(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), RequestParams, &Params{req:r})
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func CurrentParams(req *http.Request) *Params {
	return req.Context().Value(RequestParams).(*Params)
}

// Params return path variables and query variables
type Params struct {
	req         *http.Request
	queryParams map[string]string
}

func (p *Params) PathParam(key string) string {
	return chi.URLParam(p.req, key)
}

func (p *Params) QueryParam(key string) string {
	if p.queryParams == nil {
		p.queryParams = map[string]string{}
		rq := p.req.URL.RawQuery
		if rq != "" {
			for _, pair := range strings.Split(rq, "&") {
				kv := strings.Split(pair, "=")
				if uv, err := url.QueryUnescape(kv[1]); err != nil {
					panic(err)
				}else {
					p.queryParams[kv[0]] = uv
				}
			}
		}
	}
	return p.queryParams[key]
}

