package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"github.com/tinchourteaga/go-grpc-api-gateway/pkg/auth/pb"
	"github.com/tinchourteaga/go-grpc-api-gateway/pkg/auth/routes"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AuthServiceClient struct {
	Client pb.AuthServiceClient
}

func NewAuthServiceClient() pb.AuthServiceClient {
	connection, err := grpc.Dial(viper.GetString("AUTH_SVC_URL"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Error().Msg("could not establish connection: " + err.Error())
	}

	return pb.NewAuthServiceClient(connection)
}

func (svc *AuthServiceClient) Register(ctx *gin.Context) {
	routes.Register(ctx, svc.Client)
}

func (svc *AuthServiceClient) Login(ctx *gin.Context) {
	routes.Login(ctx, svc.Client)
}
