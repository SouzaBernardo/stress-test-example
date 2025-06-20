package usecases

import (
	"fmt"
	"net/http"

	"time"

	"github.com/SouzaBernardo/sample-api/pkg/metrics"
)

func ProcessPayment(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "unprocessable entity", http.StatusUnprocessableEntity)
		return
	}
	fmt.Println("Processing payment...")
	get := r.Header.Get("Status")
	if get != "success" {
		time.Sleep(20 * time.Second)
		fmt.Fprintf(w, "Payment processing failed")
		metrics.PaymentCounter.WithLabelValues("failed").Observe(1)
		return
	}
	time.Sleep(2 * time.Second)
	fmt.Fprintf(w, "Payment processed successfully")
	metrics.PaymentCounter.WithLabelValues("success").Observe(1)
}
