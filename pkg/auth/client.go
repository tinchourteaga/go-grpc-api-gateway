package auth

import (
	"github.com/rs/zerolog/log"
	"github.com/tinchourteaga/go-grpc-api-gateway/pkg/auth/pb"
	"github.com/tinchourteaga/go-grpc-api-gateway/pkg/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AuthServiceClient struct {
	Client pb.AuthServiceClient
}

func NewAuthServiceClient(c *config.Config) pb.AuthServiceClient {
	connection, err := grpc.Dial(c.AuthSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Error().Msg("could not establish connection: " + err.Error())
	}

	return pb.NewAuthServiceClient(connection)
}
