package grafanaproxy

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"regexp"
	"strings"

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

	isAllowed, err := isUpstreamURLAllowed(r)
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

func isUpstreamURLAllowed(r *http.Request) (bool, error) {
	// some static paths we need to allow
	alwaysAllowed := []string{
		"/api/annotations",
		"/api/prometheus/grafana/api/v1/rules",
	}
	for _, aa := range alwaysAllowed {
		if r.URL.Path == aa {
			return true, nil
		}
	}

	alwaysAllowedPrefixes := []string{
		"/public/",
	}
	for _, aa := range alwaysAllowedPrefixes {
		if strings.HasPrefix(r.URL.Path, aa) {
			return true, nil
		}
	}
	// queries, we need to make sure not any query is allowed
	if r.URL.Path == "/api/ds/query" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return false, errors.Wrap(err, "read request body")
		}

		type queryDatasourceReqeust struct {
			UID  string `json:"uid"`
			Type string `json:"type"`
		}

		type queryRequest struct {
			RefID         string                 `json:"refId"`
			Datasource    queryDatasourceReqeust `json:"datasource"`
			RawSQL        string                 `json:"rawSql"`
			Format        string                 `json:"format"`
			DatasourceID  int                    `json:"datasourceId"`
			IntervalMS    int64                  `json:"intervalMs"`
			MaxDataPoints int                    `json:"maxDataPoints"`
		}

		type queriesRequest struct {
			Queries []queryRequest `json:"queries"`
		}

		query := queriesRequest{}
		if err := json.Unmarshal(body, &query); err != nil {
			return false, errors.Wrap(err, "unmarshal query request")
		}

		unknownQuery := false
		for _, q := range query.Queries {
			/*
				SELECT
					"time" AS "time",
					pop_ping_latency_ms
					FROM dishy_data
					WHERE
					$__timeFilter("time")
					ORDER BY 1
			*/

			// first, make sure we have only a single statement
			if strings.Count(q.RawSQL, ";") > 1 {
				unknownQuery = true
			}

			// now, make sure we are only selecting from our table
			if !strings.Contains(q.RawSQL, "FROM dishy_data") && !strings.Contains(q.RawSQL, "FROM dishy_speed") {
				unknownQuery = true
			}

			if unknownQuery {
				fmt.Printf("unknown query: %s\n", q.RawSQL)
			}
		}

		if unknownQuery {
			return false, nil
		}

		r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		return true, nil
	}
	// dashboards: /d/<id>/default-dashboard
	// dashboards api: /api/dashboards/uid/<id>

	dashboardRegexes := []string{
		`(?:\/d\/)(?P<DishyID>.*)(?:\/default-dashboard)`,
		`(?:\/d\/)(?P<DishyID>.*)(?:\/the-data)`,
		`(?:\/api\/dashboards\/uid\/)(?P<DishyID>.*)`,
	}

	for _, dashdashboardRegex := range dashboardRegexes {
		rg := regexp.MustCompile(dashdashboardRegex)
		regexMatch := rg.FindAllStringSubmatch(r.URL.Path, -1)
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
