package prometheus

import (
	"net/http"
	"os"

	"github.com/dpjungmin/jellypi-server/utils/logger"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// StartExposition will start exposing Prometheus metrics
func StartExposition() {
	logger.Info("[PROMETHEUSE]::[STARTING_UP]")

	server := &http.Server{Addr: ":2112"}

	http.Handle("/metrics", promhttp.Handler())

	if err := server.ListenAndServe(); err != nil {
		logger.Error("[PROMETHEUSE]::[EXPOSITION_FAILURE]", err)
		os.Exit(1)
	}
}
