package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"github.com/tinchourteaga/go-grpc-api-gateway/pkg/auth"
	"github.com/tinchourteaga/go-grpc-api-gateway/pkg/config"
	"github.com/tinchourteaga/go-grpc-api-gateway/pkg/order"
	"github.com/tinchourteaga/go-grpc-api-gateway/pkg/product"
)

func main() {
	err := config.Load()
	if err != nil {
		log.Error().Msg("config loading: " + err.Error())
	}

	router := gin.Default()

	authSvc := auth.MapRoutes(router)
	order.MapRoutes(router, authSvc)
	product.MapRoutes(router, authSvc)

	router.Run(viper.GetString("PORT"))
}
