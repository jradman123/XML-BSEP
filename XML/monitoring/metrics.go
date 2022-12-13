package monitoring

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"net/http"
	"strconv"
)

type MetricsMiddleware struct {
	status200 *prometheus.CounterVec
	status500 *prometheus.CounterVec
}

func NewMetricsMiddleware() *MetricsMiddleware {
	status200 := promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "status200",
		Help: "The total number of requests with status 200 ",
	}, []string{"method", "path", "statuscode"})
	status500 := promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "status500",
		Help: "The total number of requests with status 500",
	}, []string{"method", "path", "statuscode"})
	return &MetricsMiddleware{
		status200: status200,
		status500: status500,
	}
}

// Metrics middleware to collect metrics from http requests
func (lm *MetricsMiddleware) Metrics(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		wi := &responseWriterInterceptor{
			statusCode:     http.StatusOK,
			ResponseWriter: w,
		}
		next.ServeHTTP(wi, r)

		if wi.statusCode == 200 {
			lm.status200.With(prometheus.Labels{"method": r.Method, "path": r.RequestURI, "statuscode": strconv.Itoa(wi.statusCode)}).Inc()
		} else if wi.statusCode == 500 {
			lm.status500.With(prometheus.Labels{"method": r.Method, "path": r.RequestURI, "statuscode": strconv.Itoa(wi.statusCode)}).Inc()
		}
	})
}

// responseWriterInterceptor is a simple wrapper to intercept set data on a
// ResponseWriter.
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
