package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
)

var HttpDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
	Name:    "http_duration_seconds",
	Help:    "Http request duration in seconds",
	Buckets: prometheus.DefBuckets,
}, []string{"handler"})

func init() {
	prometheus.MustRegister(HttpDuration)
}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		duration := prometheus.NewTimer(HttpDuration.WithLabelValues(r.URL.Path))
		next.ServeHTTP(w, r)
		duration.ObserveDuration()
	})
}
