package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"time"
)

var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "example_counter_app_processed_total",
		Help: "The total number of counter events",
	})
)


func countMetrics() {
	for {
		opsProcessed.Add(10)
		time.Sleep(2 * time.Second)
	}
}

func main() {

	go countMetrics()

	http.Handle("/metrics", promhttp.Handler())
	fmt.Println("prometheus-counter-example started on 0.0.0.0:9000")
	http.ListenAndServe(":9000", nil)
}
