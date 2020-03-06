package monitoring

import (
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
	)
	prometheus.Register(summaryVec)
	return summaryVec
}

func Monitor(serviceName,routeName,signature string) func(http.Handler) http.Handler{
	summaryVec :=BuildSummaryVec(serviceName,routeName,signature)

	return func(next http.Hanlder) http.Handler{
		return http.Handlerfun(func (w http.ResponseWriter,r *http.Request){
			start:=time.Now()
			next.ServiceHTTP(w,r)
			stop:=time.Now()

			summaryVec.WithLabelValues("duration").Observe(duration.Seconds())

			size,err := strconv.Atoi(w.Header().Get("Content-Length"))
			if err!=nil{
				fmt.Println("error in monitoring")
			}

		})
	}
}
