package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	service "github.com/iacopoghilardi/mynance-service-api/api/v1/services"
	"github.com/iacopoghilardi/mynance-service-api/internal/utils"
)

func AuthMiddleware(authService *service.AuthService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader("Authorization")

		if authorizationHeader == "" {
			ctx.JSON(http.StatusUnauthorized, utils.GenerateErrorResponse("Unauthorized"))
			ctx.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authorizationHeader, "Bearer ")
		if tokenString == authorizationHeader {
			ctx.JSON(http.StatusUnauthorized, utils.GenerateErrorResponse("Invalid token format"))
			ctx.Abort()
			return
		}

		user, err := authService.VerifyToken(tokenString)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, utils.GenerateErrorResponse("Invalid token"))
			ctx.Abort()
			return
		}

		ctx.Set("user", user)
		ctx.Next()
	}
}
