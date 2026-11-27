package router

import (
	"net/http"

	"github.com/PADIGUNDLA/cloud-native-devops-platform/internal/handlers"
	"github.com/PADIGUNDLA/cloud-native-devops-platform/internal/middleware"
	"github.com/PADIGUNDLA/cloud-native-devops-platform/internal/metrics"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func RegisterRoutes() {

	metrics.Init()

	mux := http.NewServeMux()

	// Static Files
	fs := http.FileServer(http.Dir("internal/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	// Dashboard
	mux.HandleFunc("/", handlers.Dashboard)

	// Health
	mux.HandleFunc("/health", handlers.Health)
	mux.HandleFunc("/ready", handlers.Health)
	mux.HandleFunc("/live", handlers.Health)

	// Metrics
	mux.Handle("/metrics", promhttp.Handler())

	// APIs
	mux.HandleFunc("/api/info", handlers.ApplicationInfo)
	mux.HandleFunc("/api/build", handlers.BuildInfo)
	mux.HandleFunc("/api/runtime", handlers.RuntimeInfo)
	mux.HandleFunc("/api/kubernetes", handlers.KubernetesInfo)

	// Middleware chain
	http.Handle("/", middleware.Logging(
		middleware.Metrics(mux),
	))
}
