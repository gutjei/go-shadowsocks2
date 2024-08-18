package metrics

import (
	"github.com/google/martian/log"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func Start(listenAddr string) {
	http.Handle("/metrics", promhttp.Handler())
	go func() {
		err := http.ListenAndServe(listenAddr, nil)
		if err != nil {
			log.Errorf("Failed to start metrics server: %s", err)
		}
	}()
}

var (
	ReceiveBytesTotal = promauto.NewCounter(prometheus.CounterOpts{
		Name: "shadowsocks_receive_bytes_total",
		Help: "Send bytes total",
	})
	TransmitBytesTotal = promauto.NewCounter(prometheus.CounterOpts{
		Name: "shadowsocks_transmit_bytes_total",
		Help: "Transmit bytes total",
	})
)
