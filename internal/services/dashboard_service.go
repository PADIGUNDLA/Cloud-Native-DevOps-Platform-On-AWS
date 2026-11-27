package services

import (
	"os"
	"time"

	"github.com/PADIGUNDLA/cloud-native-devops-platform/internal/config"
	"github.com/PADIGUNDLA/cloud-native-devops-platform/internal/models"
)

var startTime = time.Now()

func GetDashboardData(cfg config.Config) models.DashboardData {

	hostname, _ := os.Hostname()

	return models.DashboardData{
		AppName:     cfg.AppName,
		Version:     cfg.Version,
		Environment: cfg.Environment,
		Hostname:    hostname,
		CurrentTime: time.Now().Format("02 Jan 2006 15:04:05"),
		Uptime:      time.Since(startTime).Round(time.Second).String(),
	}
}
