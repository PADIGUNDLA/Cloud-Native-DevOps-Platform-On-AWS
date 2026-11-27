package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/PADIGUNDLA/cloud-native-devops-platform/internal/services"
)

func KubernetesInfo(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(services.GetKubernetesInfo())
}
