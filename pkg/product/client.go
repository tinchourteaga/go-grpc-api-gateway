package product

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"github.com/tinchourteaga/go-grpc-api-gateway/pkg/product/pb"
	"github.com/tinchourteaga/go-grpc-api-gateway/pkg/product/routes"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ProductServiceClient struct {
	Client pb.ProductServiceClient
}

func NewProductServiceClient() pb.ProductServiceClient {
	connection, err := grpc.Dial(viper.GetString("PRODUCT_SVC_URL"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Error().Msg("could not establish connection: " + err.Error())
	}

	return pb.NewProductServiceClient(connection)
}

func (svc *ProductServiceClient) CreateProduct(ctx *gin.Context) {
	routes.CreateProduct(ctx, svc.Client)
}

func (svc *ProductServiceClient) FindOne(ctx *gin.Context) {
	routes.FindOne(ctx, svc.Client)
}
