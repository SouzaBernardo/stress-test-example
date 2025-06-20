package routes

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Metrics(mux *http.ServeMux) {
	mux.Handle("GET /metrics", promhttp.Handler())
}
