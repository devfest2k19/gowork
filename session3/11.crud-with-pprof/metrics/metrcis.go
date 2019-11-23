package metrics

import (
	"fmt"
	"github.com/pickme-go/log"
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

type Metricer interface {
	Count(labels []string)
	CountLatency(time time.Time, labels []string)
}

type serverLatencyCounter struct {
	serviceLatency *prometheus.SummaryVec
}

func InitServiceLatencyCounter(namespace string, subsystem string) Metricer {

	sc := &serverLatencyCounter{
		serviceLatency: prometheus.NewSummaryVec(prometheus.SummaryOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      `request_latency`,
			Help:      `http_request_latency_milliseconds`,
		}, []string{"method"}),
	}

	err := prometheus.Register(sc.serviceLatency)
	if err != nil {
		log.Fatal(log.WithPrefix(`metrics.service_latency.go`, fmt.Sprintf(`metrics can not be registered due to : %v`, err)))
	}

	return sc
}

func (slc *serverLatencyCounter) Count(labels []string) {
	panic("implement me")
}

func (slc *serverLatencyCounter) CountLatency(begin time.Time, labels []string) {
	slc.serviceLatency.WithLabelValues(labels...).Observe(float64(time.Since(begin).Nanoseconds()) / 1e6)
}
