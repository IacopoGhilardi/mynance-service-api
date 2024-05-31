package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	prefix := r.Group("/api/v1")

	SetupUserRoutes(prefix)
}
