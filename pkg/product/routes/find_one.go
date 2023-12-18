package routes

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tinchourteaga/go-grpc-api-gateway/pkg/product/pb"
)

func FindOne(ctx *gin.Context, productSvcClient pb.ProductServiceClient) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("invalid id"))
		return
	}

	pbFindOneReq := pb.FindOneRequest{
		Id: int64(id),
	}

	res, err := productSvcClient.FindOne(context.Background(), &pbFindOneReq)
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(int(res.Status), &res)
}
