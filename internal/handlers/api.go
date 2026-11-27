package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/PADIGUNDLA/cloud-native-devops-platform/internal/config"
	"github.com/PADIGUNDLA/cloud-native-devops-platform/internal/services"
)

func ApplicationInfo(w http.ResponseWriter, r *http.Request) {

	cfg := config.Load()

	data := services.GetDashboardData(cfg)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(data)
}
