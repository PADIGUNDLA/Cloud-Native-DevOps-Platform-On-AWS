package handlers

import (
	"html/template"
	"net/http"

	"github.com/PADIGUNDLA/cloud-native-devops-platform/internal/config"
	"github.com/PADIGUNDLA/cloud-native-devops-platform/internal/services"
)

func Dashboard(w http.ResponseWriter, r *http.Request) {

	cfg := config.Load()

	data := services.GetDashboardData(cfg)

	tmpl := template.Must(
		template.ParseFiles("internal/templates/index.html"),
	)

	tmpl.Execute(w, data)
}
