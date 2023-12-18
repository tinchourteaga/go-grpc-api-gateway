package order

import (
	"github.com/gin-gonic/gin"
	"github.com/tinchourteaga/go-grpc-api-gateway/pkg/auth"
)

func MapRoutes(r *gin.Engine, authSvc *auth.AuthServiceClient) {
	authMW := auth.NewAuthMiddleware(authSvc)
	svc := &OrderServiceClient{
		Client: NewOrderServiceClient(),
	}

	routes := r.Group("/auth")
	routes.Use(authMW.Authenticate)
	routes.POST("/", svc.CreateOrder)
}
