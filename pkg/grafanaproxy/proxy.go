package grafanaproxy

import (
	"context"
	"fmt"
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

	proxy := httputil.NewSingleHostReverseProxy(url)

	fmt.Printf("r.url.path = %s, url.path = %s\n", r.URL.Path, url.Path)
	r.URL.Host = url.Host
	r.URL.Scheme = url.Scheme
	r.Header.Set("X-Forwarded-Host", r.Header.Get("Host"))
	r.Host = url.Host

	logger.Info("proxying request",
		zap.String("upstreamEndpoint", url.Host),
		zap.String("requestedPath", r.URL.Path),
		zap.String("scheme", url.Scheme),
		zap.String("x-forwarded-host", r.Header.Get("Host")))

	proxy.ServeHTTP(w, r)

}

func grafanaEndpointForRequest(r *http.Request) (string, error) {
	// if the path is a root dashboard path, we'll rewrite it
	// for all other internal grafana paths (CSS, images, etc)
	// we keep them the same

	// if r.URL.Path == "/23txAzw1aH0rlpXYhOoleMHC2Fe" {
	// 	return "http://grafana:3000/d/apzwaMb7z/sample?orgId=1", nil
	// }

	return "http://grafana:3000", nil
}
