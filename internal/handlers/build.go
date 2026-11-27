package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/PADIGUNDLA/cloud-native-devops-platform/internal/services"
)

func BuildInfo(w http.ResponseWriter, r *http.Request) {

	build := services.GetBuildInfo()

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(build)
}
