package monitoring

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

func BuildSummaryVec(serviceName, metricName, metricHelp string) *prometheus.SummaryVec {
	summaryVec := prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Namespace: serviceName,
			Name:      metricName,
			Help:      metricHelp,
		},
		[]string{"service"},
	)
	prometheus.Register(summaryVec)
	return summaryVec
}
func Monitor(serviceName, routeName, signature string) func(http.Handler) http.Handler {
	summaryVec := BuildSummaryVec(serviceName, routeName, signature)
	fmt.Println("inside monitor")
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			fmt.Println("before monitor")
			next.ServeHTTP(w, r)
			fmt.Println("after monitor")
			duration := time.Since(start)

			summaryVec.WithLabelValues("duration").Observe(duration.Seconds())

			size, err := strconv.Atoi(w.Header().Get("Content-Length"))
			if err != nil {
				summaryVec.WithLabelValues("size").Observe(float64(size))
			}

		})
	}
}
