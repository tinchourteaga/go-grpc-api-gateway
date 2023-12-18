package product

import (
	"github.com/gin-gonic/gin"
	"github.com/tinchourteaga/go-grpc-api-gateway/pkg/auth"
)

func MapRoutes(r *gin.Engine, authSvc *auth.AuthServiceClient) {
	authMW := auth.NewAuthMiddleware(authSvc)
	svc := &ProductServiceClient{
		Client: NewProductServiceClient(),
	}

	rg := r.Group("/product")
	rg.Use(authMW.Authenticate)
	rg.POST("/", svc.CreateProduct)
	rg.GET("/:id", svc.FindOne)

}
