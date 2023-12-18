package auth

import (
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tinchourteaga/go-grpc-api-gateway/pkg/auth/pb"
)

type AuthMiddleware struct {
	authSvcClient *AuthServiceClient
}

func NewAuthMiddleware(svc *AuthServiceClient) AuthMiddleware {
	return AuthMiddleware{authSvcClient: svc}
}

func (am *AuthMiddleware) Authenticate(ctx *gin.Context) {
	authorization := ctx.Request.Header.Get("authorization")

	if authorization == "" {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token := strings.Split(authorization, "Bearer ")

	if len(token) < 2 {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	pbValidateReq := pb.ValidateRequest{
		Token: token[1],
	}

	res, err := am.authSvcClient.Client.Validate(context.Background(), &pbValidateReq)
	if err != nil || res.Status != http.StatusOK {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	ctx.Set("userId", res.UserId)
	ctx.Next()
}
