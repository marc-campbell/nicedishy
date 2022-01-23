package clusterproxy

import (
	"context"
	"crypto/tls"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/marc-campbell/nicedishy/pkg/logger"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func Start(ctx context.Context) {
	logger.Info("starting grafana proxy on port 3000")
	http.HandleFunc("/", handleRequestAndRedirect)

	if err := http.ListenAndServe(":3000", nil); err != nil {
		panic(err)
	}
}

func handleRequestAndRedirect(w http.ResponseWriter, r *http.Request) {
	logger.Debug("handling grafana proxy request")

	upstreamEndpoint, err := grafanaEndpointForRequest(r)
	if err != nil {
		logger.Error(errors.Wrap(err, "get cluster endpoint for request"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	url, err := url.Parse(upstreamEndpoint)
	if err != nil {
		logger.Error(errors.Wrap(err, "parse upstream uri"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	logger.Info("proxying request",
		zap.String("upstreamEndpoint", url.Host),
		zap.String("path", "/"),
		zap.String("scheme", url.Scheme),
		zap.String("x-forwarded-host", r.Header.Get("Host")))

	proxy := httputil.NewSingleHostReverseProxy(url)
	proxy.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	r.URL.Host = url.Host
	r.URL.Scheme = url.Scheme
	r.Header.Set("X-Forwarded-Host", r.Header.Get("Host"))
	r.Host = url.Host
	r.URL.Path = "/"

	proxy.ServeHTTP(w, r)

}

func grafanaEndpointForRequest(r *http.Request) (string, error) {
	return "https://grafana:3000/d/apzwaMb7z/sample?orgId=1&from=now-6h&to=now", nil
}
