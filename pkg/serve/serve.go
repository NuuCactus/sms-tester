package serve

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/nuucactus/sms-tester/router"
	"github.com/spf13/cobra"
)

// RunServe the main event loop for the service
func RunServe() func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {

		urlRestAPI, err := cmd.Flags().GetString("rest-api-url")
		if err != nil {
			log.Println("Missing rest-api-url")
			os.Exit(1)
		}

		urlMetricsAPI, err := cmd.Flags().GetString("metrics-api-url")
		if err != nil {
			log.Println("Missing rest-api-url")
			os.Exit(1)
		}

		uRestAPI, err := url.Parse(urlRestAPI)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}

		uMetricsAPI, err := url.Parse(urlMetricsAPI)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}

		rMon := router.NewRouterForMetricsAPI()
		rAPI := router.NewRouterForRestAPI()

		srvMetricsAPI := &http.Server{
			Addr: fmt.Sprintf("%s:%s", uMetricsAPI.Hostname(), uMetricsAPI.Port()),
			// Good practice to set timeouts to avoid Slowloris attacks.
			WriteTimeout: time.Second * 15,
			ReadTimeout:  time.Second * 15,
			IdleTimeout:  time.Second * 60,
			Handler:      rMon, // Pass our instance of gorilla/mux in.
		}

		srvRestAPI := &http.Server{
			Addr: fmt.Sprintf("%s:%s", uRestAPI.Hostname(), uRestAPI.Port()),
			// Good practice to set timeouts to avoid Slowloris attacks.
			WriteTimeout: time.Second * 15,
			ReadTimeout:  time.Second * 15,
			IdleTimeout:  time.Second * 60,
			Handler:      rAPI, // Pass our instance of gorilla/mux in.
		}

		// Run our server in a goroutine so that it doesn't block.
		go func() {
			log.Printf("Metrics listening for %s connections on %s:%s", uMetricsAPI.Scheme, uMetricsAPI.Hostname(), uMetricsAPI.Port())
			if err := srvMetricsAPI.ListenAndServe(); err != nil {
				log.Fatal(err)
			}
		}()

		go func() {
			log.Printf("Rest API listening for %s connections on %s:%s", uRestAPI.Scheme, uRestAPI.Hostname(), uRestAPI.Port())
			if err := srvRestAPI.ListenAndServe(); err != nil {
				log.Fatal(err)
			}
		}()

		c := make(chan os.Signal, 1)
		// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
		// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
		signal.Notify(c, os.Interrupt)

		// Block until we receive our signal.
		<-c

		// Create a deadline to wait for.
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
		defer cancel()
		// Doesn't block if no connections, but will otherwise wait
		// until the timeout deadline.
		srvRestAPI.Shutdown(ctx)
		srvMetricsAPI.Shutdown(ctx)
		// Optionally, you could run srv.Shutdown in a goroutine and block on
		// <-ctx.Done() if your application should wait for other services
		// to finalize based on context cancellation.
		log.Println("shutting down")
		os.Exit(0)
	}
}
