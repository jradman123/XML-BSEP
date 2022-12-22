package monitoring

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"net/http"
	"strconv"
	"strings"
)

type MetricsMiddleware struct {
	totalRequests      *prometheus.CounterVec
	durationOfRequests *prometheus.HistogramVec
}

func NewMetricsMiddleware() *MetricsMiddleware {
	totalRequests := promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "The total number of requests.",
	}, []string{"method", "path", "statusCode", "service"})
	durationOfRequests := promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name: "http_response_time_seconds",
		Help: "Duration of HTTP requests.",
	}, []string{"path", "service"})
	return &MetricsMiddleware{
		totalRequests:      totalRequests,
		durationOfRequests: durationOfRequests,
	}
}

// Metrics middleware to collect metrics from http requests
func (lm *MetricsMiddleware) Metrics(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		timer := prometheus.NewTimer(lm.durationOfRequests.With(prometheus.Labels{"path": r.RequestURI, "service": getServiceName(r.RequestURI)}))

		wi := &responseWriterInterceptor{
			statusCode:     http.StatusOK,
			ResponseWriter: w,
		}

		next.ServeHTTP(wi, r)
		lm.totalRequests.With(prometheus.Labels{"method": r.Method, "path": r.RequestURI, "statusCode": strconv.Itoa(wi.statusCode), "service": getServiceName(r.RequestURI)}).Inc()
		timer.ObserveDuration()
	})
}

type responseWriterInterceptor struct {
	http.ResponseWriter
	statusCode int
}

func (w *responseWriterInterceptor) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func (w *responseWriterInterceptor) Write(p []byte) (int, error) {
	return w.ResponseWriter.Write(p)
}

func getServiceName(path string) (serviceName string) {
	partOfPath := strings.Split(path, "/")
	fmt.Println(partOfPath[1])
	if strings.EqualFold(partOfPath[1], "connection") {
		serviceName = "connection_service"
	} else if strings.EqualFold(partOfPath[1], "job_offer") || strings.EqualFold(partOfPath[1], "post") {
		serviceName = "post_service"
	} else if strings.EqualFold(partOfPath[1], "messages") || strings.EqualFold(partOfPath[1], "notification") {
		serviceName = "message_service"
	} else if strings.EqualFold(partOfPath[1], "users") && !strings.EqualFold(partOfPath[2], "auth") && !strings.EqualFold(partOfPath[2], "login") && !strings.EqualFold(partOfPath[3], "feed") {
		serviceName = "user_service"
	} else if strings.EqualFold(partOfPath[1], "users") && (strings.EqualFold(partOfPath[2], "auth") || strings.EqualFold(partOfPath[2], "login") || strings.EqualFold(partOfPath[3], "feed")) {
		serviceName = "api_gateway"
	} else {
		serviceName = "api_gateway"
	}
	return serviceName
}
