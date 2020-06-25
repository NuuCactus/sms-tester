package router

import (
	"github.com/nuucactus/sms-tester/endpoints/restapi"
	"github.com/nuucactus/sms-worker/endpoints/metrics"

	"github.com/gorilla/mux"
)

// NewRouterForMetricsAPI creates a mux router with definitions for metrics endpoints
func NewRouterForMetricsAPI() (r *mux.Router) {
	r = mux.NewRouter()
	r.HandleFunc("/metrics", metrics.GetMetrics).Methods("GET")
	return r
}

// NewRouterForRestAPI creates a mux router with definitions for rest api endpoints
func NewRouterForRestAPI() (r *mux.Router) {
	r = mux.NewRouter()
	r.HandleFunc("/sms", restapi.PostSMS).Methods("POST")
	return r
}
