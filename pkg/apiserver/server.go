package apiserver

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
	"github.com/gorilla/mux"
	"github.com/marc-campbell/nicedishy/pkg/handlers"
)

func Start() {
	log.Printf("nicedishy api version %s\n", os.Getenv("VERSION"))

	// initialize the oauth2 server
	manager := manage.NewDefaultManager()
	manager.MustTokenStorage(store.NewMemoryTokenStore()) // TODO postgres/

	clientStore := store.NewClientStore()
	clientStore.Set("000000", &models.Client{
		ID:     "000000",
		Secret: "999999",
		Domain: "http://localhost",
	})
	manager.MapClientStorage(clientStore)

	oauthServer := server.NewDefaultServer(manager)
	oauthServer.SetAllowGetAccessRequest(true)
	oauthServer.SetClientInfoHandler(server.ClientFormHandler)
	oauthServer.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Println("Internal Error:", err.Error())
		return
	})
	oauthServer.SetResponseErrorHandler(func(re *errors.Response) {
		log.Println("Response Error:", re.Error.Error())
	})

	r := mux.NewRouter()

	r.Methods("OPTIONS").HandlerFunc(handlers.CORS)

	/**********************************************************************
	* Unauthenticated routes
	**********************************************************************/

	r.HandleFunc("/healthz", handlers.Healthz)
	r.HandleFunc("/api/v1/login", handlers.Login)
	r.Path("/api/v1/login/callback").Methods("POST").HandlerFunc(handlers.LoginCallback)

	// oauth server routes (auth is handled by the package)
	r.HandleFunc("/authorize", func(w http.ResponseWriter, r *http.Request) {
		err := oauthServer.HandleAuthorizeRequest(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	})
	r.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
		oauthServer.HandleTokenRequest(w, r)
	})

	/**********************************************************************
	* Static routes
	**********************************************************************/

	/**********************************************************************
	* Authenticated routes
	**********************************************************************/
	sessionAuthQuietRouter := r.PathPrefix("").Subrouter()
	sessionAuthQuietRouter.Use(handlers.RequireValidSessionQuietMiddleware)
	sessionAuthQuietRouter.Path("/api/v1/dishies").Methods("GET").HandlerFunc(handlers.GetDishies)

	srv := &http.Server{
		Handler: r,
		Addr:    ":3000",
	}

	fmt.Printf("Starting nicedishy api on port %d...\n", 3000)

	log.Fatal(srv.ListenAndServe())
}
