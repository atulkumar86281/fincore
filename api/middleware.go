package api

import (
	"errors"
	"github/atulkumar0001/Bank/token"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const(
	authorizationHeaderKey = "authorization"
	authorizationHeaderType = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

func authMiddleware(tokenMaker token.Maker)(gin.HandlerFunc){

	return func(ctx *gin.Context){
		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)

		if len(authorizationHeader) == 0 {
			err := errors.New("Authorization header is not provided")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized,errorResponse(err))
			return
		}

		fields := strings.Fields(authorizationHeader)

		if len(fields) < 2 {
			err := errors.New("Invalid authorization header is provided")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized,errorResponse(err))
			return
		}

		authorizationType := strings.ToLower(fields[0])

		if authorizationType != authorizationHeaderType{
			err := errors.New("Unsupported authorization Type is provided")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized,errorResponse(err))
			return
		}

		accessToken := fields[1]

		tokenPayload, err := tokenMaker.VerifyToken(accessToken)

		if err != nil{
			ctx.AbortWithStatusJSON(http.StatusUnauthorized,errorResponse(err))
			return
		}

		ctx.Set(authorizationPayloadKey,tokenPayload)
	}
}