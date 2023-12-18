package auth

import (
	"github.com/gin-gonic/gin"
)

func MapRoutes(r *gin.Engine) *AuthServiceClient {
	svc := &AuthServiceClient{
		Client: NewAuthServiceClient(),
	}

	rg := r.Group("/auth")
	rg.POST("/register", svc.Register)
	rg.POST("/login", svc.Login)

	return svc
}
