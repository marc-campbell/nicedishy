package grafanaproxy

import (
	"context"
	"net/http"
	"net/http/httputil"
	"net/url"
	"regexp"

	"github.com/marc-campbell/nicedishy/pkg/logger"
	"github.com/marc-campbell/nicedishy/pkg/stores"
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

	isAllowed, err := isUpstreamURLAllowed(r.URL.Path)
	if err != nil {
		logger.Error(errors.Wrap(err, "check if upstream url is allowed"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if !isAllowed {
		logger.Infof("cowardly refusing to proxy request to %s", r.URL.Path)
		w.WriteHeader(http.StatusNotFound)
		return
	}

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

func isUpstreamURLAllowed(path string) (bool, error) {
	// some static paths we need to allow
	alwaysAllowed := []string{
		"/api/ds/query",
		"/api/annotations",
		"/api/prometheus/grafana/api/v1/rules",
	}
	for _, aa := range alwaysAllowed {
		if path == aa {
			return true, nil
		}
	}

	// dashboards: /d/<id>/default-dashboard
	// dashboards api: /api/dashboards/uid/<id>

	dashboardRegexes := []string{
		`(?:\/d\/)(?P<DishyID>.*)(?:\/default-dashboard)`,
		`(?:\/api\/dashboards\/uid\/)(?P<DishyID>.*)`,
	}

	for _, dashdashboardRegex := range dashboardRegexes {
		r := regexp.MustCompile(dashdashboardRegex)
		regexMatch := r.FindAllStringSubmatch(path, -1)
		for _, result := range regexMatch {
			maybeDishyID := result[1]
			maybeDishy, err := stores.GetStore().GetDishy(context.TODO(), maybeDishyID)
			if err != nil {
				return false, errors.Wrap(err, "get dishy")
			}
			if maybeDishy != nil {
				return true, nil
			}

			// don't return false, let the end handle it
		}
	}

	return false, nil
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
