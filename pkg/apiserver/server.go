package apiserver

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"

	"github.com/gorilla/mux"
	"github.com/marc-campbell/nicedishy/pkg/handlers"
)

func Start() {
	log.Printf("nicedishy api version %s\n", os.Getenv("VERSION"))

	r := mux.NewRouter()

	r.Methods("OPTIONS").HandlerFunc(handlers.CORS)

	/**********************************************************************
	* Unauthenticated routes
	**********************************************************************/

	r.HandleFunc("/healthz", handlers.Healthz)
	r.Path("/api/v1/subscribe").Methods("POST").HandlerFunc(handlers.Subscribe)

	/**********************************************************************
	* Static routes
	**********************************************************************/

	/**********************************************************************
	* Authenticated routes
	**********************************************************************/

	/**********************************************************************
	* Internal API
	**********************************************************************/
	internalQuietRouter := r.PathPrefix("/api/internal").Subrouter()
	internalQuietRouter.Use(handlers.RequireValidInternalAuthQuietMiddleware)
	internalQuietRouter.Path("/dishy/{id}/dashboard").Methods("PUT").HandlerFunc(handlers.DeployLatestDishyDashboard)
	internalQuietRouter.Path("/public/dashboards").Methods("PUT").HandlerFunc(handlers.DeployLatestPublicDashboards)

	/**********************************************************************
	* Grafana
	**********************************************************************/
	u, err := url.Parse(os.Getenv("GRAFANA_ENDPOINT"))
	if err != nil {
		panic(err)
	}
	upstream := httputil.NewSingleHostReverseProxy(u)
	grafanaProxy := func(upstream *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
		return func(w http.ResponseWriter, r *http.Request) {
			r.Header.Set("X-Forwarded-Host", r.Header.Get("Host"))
			upstream.ServeHTTP(w, r)
		}
	}(upstream)
	r.PathPrefix("/proxy/grafana").HandlerFunc(grafanaProxy)

	tokenAuthQuietRouter := r.PathPrefix("").Subrouter()
	tokenAuthQuietRouter.Use(handlers.RequireValidTokenQuietMiddleware)
	tokenAuthQuietRouter.Path("/api/v1/stats").Methods("POST").HandlerFunc(handlers.StoreData)
	tokenAuthQuietRouter.Path("/api/v1/speed").Methods("POST").HandlerFunc(handlers.StoreSpeed)

	srv := &http.Server{
		Handler: r,
		Addr:    ":3000",
	}

	fmt.Printf("Starting nicedishy-api on port %d...\n", 3000)

	log.Fatal(srv.ListenAndServe())
}
