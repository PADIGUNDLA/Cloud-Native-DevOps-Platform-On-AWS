package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/PADIGUNDLA/cloud-native-devops-platform/internal/services"
)

func RuntimeInfo(w http.ResponseWriter, r *http.Request) {

	info := services.GetRuntimeInfo()

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(info)
}
