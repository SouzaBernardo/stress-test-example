package metrics

import "github.com/prometheus/client_golang/prometheus"

var PaymentCounter = prometheus.NewHistogramVec(prometheus.HistogramOpts{
	Name:    "payment_counter",
	Help:    "Number of payments",
	Buckets: prometheus.DefBuckets,
}, []string{"status"})

func init() {
	prometheus.MustRegister(PaymentCounter)
}
