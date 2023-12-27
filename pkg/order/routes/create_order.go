package routes

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tinchourteaga/go-grpc-api-gateway/pkg/order/pb"
)

type CreateOrderRequestBody struct {
	ProductId int64 `json:"product_id"`
	Quantity  int64 `json:"quantity"`
}

func CreateOrder(ctx *gin.Context, orderSvcClient pb.OrderServiceClient) {
	body := CreateOrderRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	userId, _ := ctx.Get("userId")
	userId, err := strconv.Atoi(userId.(string))
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	pbCreateOrderReq := pb.CreateRequest{
		ProductId: body.ProductId,
		Quantity:  body.Quantity,
		UserId:    int64(userId.(int)),
	}

	res, err := orderSvcClient.CreateOrder(context.Background(), &pbCreateOrderReq)
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
