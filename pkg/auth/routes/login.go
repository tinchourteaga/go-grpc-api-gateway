package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tinchourteaga/go-grpc-api-gateway/pkg/auth/pb"
)

type LoginRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(ctx *gin.Context, authSvcClient pb.AuthServiceClient) {
	body := LoginRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	pbLoginReq := pb.LoginRequest{
		Email:    body.Email,
		Password: body.Password,
	}

	res, err := authSvcClient.Login(context.Background(), &pbLoginReq)
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(int(res.Status), &res)
}
