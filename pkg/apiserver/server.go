package apiserver

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/marc-campbell/nicedishy/pkg/handlers"
	"github.com/marc-campbell/nicedishy/pkg/stores"
	"github.com/nats-io/nats.go"
)

func Start() {
	log.Printf("nicedishy api version %s\n", os.Getenv("VERSION"))

	r := mux.NewRouter()

	r.Methods("OPTIONS").HandlerFunc(handlers.CORS)

	/**********************************************************************
	* Unauthenticated routes
	**********************************************************************/

	r.HandleFunc("/healthz", handlers.Healthz)
	r.HandleFunc("/api/v1/login", handlers.Login)
	r.Path("/api/v1/login/callback").Methods("POST").HandlerFunc(handlers.LoginCallback)
	r.Path("/api/v1/subscribe").Methods("POST").HandlerFunc(handlers.Subscribe)
	r.Path("/api/v1/stats/public/stream").Methods("GET").HandlerFunc(handlers.StreamPublicStats)

	/**********************************************************************
	* Static routes
	**********************************************************************/

	/**********************************************************************
	* Authenticated routes
	**********************************************************************/
	sessionAuthQuietRouter := r.PathPrefix("").Subrouter()
	sessionAuthQuietRouter.Use(handlers.RequireValidSessionQuietMiddleware)
	sessionAuthQuietRouter.Path("/api/v1/nonce").Methods("GET").HandlerFunc(handlers.CreateNonce)
	sessionAuthQuietRouter.Path("/api/v1/dishy").Methods("POST").HandlerFunc(handlers.CreateDishy)
	sessionAuthQuietRouter.Path("/api/v1/dishies").Methods("GET").HandlerFunc(handlers.ListDishies)
	sessionAuthQuietRouter.Path("/api/v1/dishy/{id}/token").Methods("GET").HandlerFunc(handlers.GetDishyToken)
	sessionAuthQuietRouter.Path("/api/v1/dishy/{id}").Methods("DELETE").HandlerFunc(handlers.DeleteDishy)

	/**********************************************************************
	* Nonce routes
	**********************************************************************/
	nonceAuthQuietRouter := r.PathPrefix("").Subrouter()
	nonceAuthQuietRouter.Use(handlers.RequireValidNonceMiddleware)
	nonceAuthQuietRouter.Path("/api/v1/dishies/stream").Methods("GET").HandlerFunc(handlers.StreamDishies)
	nonceAuthQuietRouter.Path("/api/v1/dishy/stream").Methods("GET").HandlerFunc(handlers.StreamDishy)

	tokenAuthQuietRouter := r.PathPrefix("").Subrouter()
	tokenAuthQuietRouter.Use(handlers.RequireValidTokenQuietMiddleware)
	tokenAuthQuietRouter.Path("/api/v1/stats").Methods("POST").HandlerFunc(handlers.StoreData)
	tokenAuthQuietRouter.Path("/api/v1/speed").Methods("POST").HandlerFunc(handlers.StoreSpeed)

	srv := &http.Server{
		Handler: r,
		Addr:    ":3000",
	}

	go func() {
		fmt.Printf("Starting public stats publisher...\n")

		opts := []nats.Option{nats.Name("public stats")}
		totalWait := 10 * time.Minute
		reconnectDelay := time.Second

		opts = append(opts, nats.ReconnectWait(reconnectDelay))
		opts = append(opts, nats.MaxReconnects(int(totalWait/reconnectDelay)))
		opts = append(opts, nats.DisconnectErrHandler(func(nc *nats.Conn, err error) {
			log.Printf("Disconnected due to:%s, will attempt reconnects for %.0fm", err, totalWait.Minutes())
		}))
		opts = append(opts, nats.ReconnectHandler(func(nc *nats.Conn) {
			log.Printf("Reconnected [%s]", nc.ConnectedUrl())
		}))
		opts = append(opts, nats.ClosedHandler(func(nc *nats.Conn) {
			log.Fatalf("Exiting: %v", nc.LastError())
		}))

		nc, err := nats.Connect("nats://nats:4222", opts...)
		if err != nil {
			log.Fatal(err)
		}

		subj := "public_stats"

		for {
			ctx := context.TODO()
			connectedDishyCount, err := stores.GetStore().GetConnectedDishyCount(ctx)
			if err != nil {
				log.Fatal(err)
				return
			}
			totalDishyCount, err := stores.GetStore().GetTotalDishyCount(ctx)
			if err != nil {
				log.Fatal(err)
				return
			}
			newDishyCount, err := stores.GetStore().GetNewDishyCount(ctx)
			if err != nil {
				log.Fatal(err)
				return
			}
			highestDownloadSpeed, err := stores.GetStore().GetHighestDownloadSpeed(ctx)
			if err != nil {
				log.Fatal(err)
				return
			}
			averageDownloadSpeed, err := stores.GetStore().GetAverageDownloadSpeed(ctx)
			if err != nil {
				log.Fatal(err)
				return
			}
			lowestPingTime, err := stores.GetStore().GetLowestPingTime(ctx)
			if err != nil {
				log.Fatal(err)
				return
			}
			averagePingTime, err := stores.GetStore().GetAveragePingTime(ctx)
			if err != nil {
				log.Fatal(err)
				return
			}

			message := handlers.StreamPublicStatsResponse{
				ConnectedDishyCount:  connectedDishyCount,
				AllTimeDishyCount:    totalDishyCount,
				NewDishyCount:        newDishyCount,
				HighestDownloadSpeed: highestDownloadSpeed,
				AverageDownloadSpeed: averageDownloadSpeed,
				LowestPingTime:       lowestPingTime,
				AveragePingTime:      averagePingTime,
			}

			b, err := json.Marshal(message)
			if err != nil {
				log.Fatal(err)
				return
			}

			if err := nc.Publish(subj, b); err != nil {
				log.Fatal(err)
				return
			}

			time.Sleep(time.Second)
		}
	}()

	fmt.Printf("Starting nicedishy-api on port %d...\n", 3000)

	log.Fatal(srv.ListenAndServe())
}
