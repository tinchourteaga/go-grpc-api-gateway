package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tinchourteaga/go-grpc-api-gateway/pkg/auth/pb"
)

type RegisterRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(ctx *gin.Context, authSvcClient pb.AuthServiceClient) {
	body := RegisterRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	pbRegisterReq := &pb.RegisterRequest{
		Email:    body.Email,
		Password: body.Password,
	}

	res, err := authSvcClient.Register(context.Background(), pbRegisterReq)
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
