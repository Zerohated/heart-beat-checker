package middleware

import (
	"golang.org/x/time/rate"

	"net/http"

	"github.com/gin-gonic/gin"
)

func LimiterMiddle(limiter *rate.Limiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !limiter.Allow() {
			c.AbortWithStatus(http.StatusTooManyRequests)
			return
		}
		c.Next()
	}
}
