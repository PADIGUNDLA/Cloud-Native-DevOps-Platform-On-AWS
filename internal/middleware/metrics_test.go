package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/PADIGUNDLA/cloud-native-devops-platform/internal/metrics"
	dto "github.com/prometheus/client_model/go"
)

func TestMetricsDoesNotCountScrapeRequests(t *testing.T) {
	metrics.Init()

	counter := metrics.RequestCount.WithLabelValues("GET", "/metrics", "200")
	beforeMetric := &dto.Metric{}
	if err := counter.Write(beforeMetric); err != nil {
		t.Fatal(err)
	}
	before := beforeMetric.GetCounter().GetValue()

	handler := Metrics(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	handler.ServeHTTP(httptest.NewRecorder(), httptest.NewRequest(http.MethodGet, "/metrics", nil))

	afterMetric := &dto.Metric{}
	if err := counter.Write(afterMetric); err != nil {
		t.Fatal(err)
	}
	if after := afterMetric.GetCounter().GetValue(); after != before {
		t.Errorf("metrics scrape incremented request counter: before=%v after=%v", before, after)
	}
}
