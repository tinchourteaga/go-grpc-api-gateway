package auth

import (
	"github.com/gin-gonic/gin"
)

func MapRoutes(r *gin.Engine) {
	svc := &AuthServiceClient{
		Client: NewAuthServiceClient(),
	}

	routes := r.Group("/auth")
	routes.POST("/register", svc.Register)
	routes.POST("/login", svc.Login)
}
