package handlers

import (
	"github.com/gin-gonic/gin"
	"go-demo-app/internal/utils/secrets"
	"net/http"
	"os"
	"runtime"
	"time"
)

var (
	startTime = time.Now()
	buildInfo = struct {
		BuildDate string
		Version   string
	}{
		BuildDate: os.Getenv("BUILD_DATE"),
		Version:   os.Getenv("APP_VERSION"),
	}
)

// StatsHandler handles the /stats endpoint
type StatsHandler struct{}

// NewStatsHandler creates a new StatsHandler
func NewStatsHandler() *StatsHandler {
	return &StatsHandler{}
}

// GetStats returns system statistics and application information
func (h *StatsHandler) GetStats(c *gin.Context) {
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)

	stats := gin.H{
		"uptime":      time.Since(startTime).String(),
		"environment": secrets.GetFromEnv("ENVIRONMENT", "null"),
		"build": gin.H{
			"goVersion": runtime.Version(),
			"buildDate": buildInfo.BuildDate,
			"version":   buildInfo.Version,
		},
	}

	c.JSON(http.StatusOK, stats)
}
