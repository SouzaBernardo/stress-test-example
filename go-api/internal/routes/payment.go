package routes

import (
	"net/http"

	"github.com/SouzaBernardo/sample-api/internal/usecases"
	"github.com/SouzaBernardo/sample-api/pkg/metrics"
)

func Payment(mux *http.ServeMux) {
	mux.Handle("GET /payments", metrics.Middleware(http.HandlerFunc(usecases.ProcessPayment)))
}
