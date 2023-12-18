package order

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"github.com/tinchourteaga/go-grpc-api-gateway/pkg/order/pb"
	"github.com/tinchourteaga/go-grpc-api-gateway/pkg/order/routes"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type OrderServiceClient struct {
	Client pb.OrderServiceClient
}

func NewOrderServiceClient() pb.OrderServiceClient {
	connection, err := grpc.Dial(viper.GetString("ORDER_SVC_URL"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Error().Msg("could not establish connection: " + err.Error())
		return nil
	}

	return pb.NewOrderServiceClient(connection)
}

func (svc *OrderServiceClient) CreateOrder(ctx *gin.Context) {
	routes.CreateOrder(ctx, svc.Client)
}
