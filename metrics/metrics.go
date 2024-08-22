package metrics

import (
	"github.com/google/martian/log"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"io"
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

func Copy(reader io.Reader, writer io.Writer, metric prometheus.Counter) (int64, error) {
	var err1, err2 error
	var b int64
	var n int
	buf := make([]byte, 1024)
	for {
		n, err1 = reader.Read(buf)
		if err1 != nil && err1 != io.EOF {
			return b, err1
		}
		n, err2 = writer.Write(buf[:n])
		if err2 != nil {
			metric.Add(float64(n))
			return b + int64(n), err2
		}
		b += int64(n)
		metric.Add(float64(n))

		if err1 == io.EOF {
			return b, nil
		}
	}
}
