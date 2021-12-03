package apiserver

import (
	"fmt"
	"log"
	"net/http"
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
	r.HandleFunc("/v1/login", handlers.Login)
	r.Path("/v1/login/callback").Methods("POST").HandlerFunc(handlers.LoginCallback)

	/**********************************************************************
	* Static routes
	**********************************************************************/

	/**********************************************************************
	* Authenticated routes
	**********************************************************************/
	sessionAuthQuietRouter := r.PathPrefix("").Subrouter()
	sessionAuthQuietRouter.Use(handlers.RequireValidSessionQuietMiddleware)
	sessionAuthQuietRouter.Path("/v1/dishies").Methods("GET").HandlerFunc(handlers.GetDishies)
	sessionAuthQuietRouter.Path("/v1/dishy/{id}/token").Methods("GET").HandlerFunc(handlers.GetDishyToken)

	tokenAuthQuietRouter := r.PathPrefix("").Subrouter()
	tokenAuthQuietRouter.Use(handlers.RequireValidTokenQuietMiddleware)
	tokenAuthQuietRouter.Path("/v1/stats").Methods("POST").HandlerFunc(handlers.StoreData)

	srv := &http.Server{
		Handler: r,
		Addr:    ":3000",
	}

	fmt.Printf("Starting nicedishy api on port %d...\n", 3000)

	log.Fatal(srv.ListenAndServe())
}
