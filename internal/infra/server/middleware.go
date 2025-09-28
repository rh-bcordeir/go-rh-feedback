package server

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	httpRequests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_request_total",
			Help: "Number of requests",
		}, []string{"path"})

	requestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "http_request_duration_seconds",
			Help: "Duration of http request",
		}, []string{"path"})
)

func init() {
	prometheus.MustRegister(httpRequests)
	prometheus.MustRegister(requestDuration)
}

func PrometheusMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// inicia timer e incrementa contador
		timer := prometheus.NewTimer(requestDuration.WithLabelValues(r.URL.Path))
		httpRequests.WithLabelValues(r.URL.Path).Inc()

		// chama o próximo handler
		next.ServeHTTP(w, r)

		// registra a duração depois que o handler terminou
		timer.ObserveDuration()
	})
}
