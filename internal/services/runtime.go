package services

import (
	"runtime"

	"github.com/PADIGUNDLA/cloud-native-devops-platform/internal/models"
)

func GetRuntimeInfo() models.RuntimeInfo {

	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)

	return models.RuntimeInfo{
		GoVersion:  runtime.Version(),
		OS:         runtime.GOOS,
		Arch:       runtime.GOARCH,
		CPU:        runtime.NumCPU(),
		Goroutines: runtime.NumGoroutine(),
		MemoryAlloc: mem.Alloc,
	}
}
