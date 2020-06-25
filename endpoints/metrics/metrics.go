package metrics

import (
	"net/http"
)

func GetMetrics(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get /metrics\n"))
}
