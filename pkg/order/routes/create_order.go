package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tinchourteaga/go-grpc-api-gateway/pkg/order/pb"
)

type CreateOrderRequestBody struct {
	ProductId int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

func CreateOrder(ctx *gin.Context, orderSvcClient pb.OrderServiceClient) {
	body := CreateOrderRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	userId, _ := ctx.Get("userId")

	pbCreateOrderReq := pb.CreateRequest{
		ProductId: int64(body.ProductId),
		Quantity:  int64(body.Quantity),
		UserId:    userId.(int64),
	}

	res, err := orderSvcClient.CreateOrder(context.Background(), &pbCreateOrderReq)
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
