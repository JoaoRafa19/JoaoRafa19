package utils

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Standard Gin
type STDLogger struct {
	logger *zap.Logger
}

func StdLogger(gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger, _ := zap.NewProduction(
			zap.Fields(
				zap.String("service", "my-service"),
				zap.String("env", "production"),
			),
		)
		defer logger.Sync()
		c.Set("logger", logger)
	}
}
