package mw

import (
	"context"
	"errors"
	"fmt"
	"lintang/go_hertz_template/biz/util/jwt"
	"net/http"
	"strings"

	"github.com/cloudwego/hertz/pkg/app"
)

// ResponseError represent the response error struct
type responseError struct {
	Message string `json:"message"`
}

const (
	AuthorizationHeaderKey  = "authorization"
	AuthorizationTypeBearer = "bearer"
	AuthorizationPayloadKey = "authorization_payload"
)

func Protected(tokenMaker jwt.JwtTokenMaker) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// pre-handle
		// ...
		// if there is no 'post-handle' logic, the 'c.Next(ctx)' can be omitted.
		authorizationHeader := c.GetHeader(AuthorizationHeaderKey)
		// if !exists {
		// 	err := errors.New("authorization header is not provided")
		// 	c.AbortWithStatusJSON(http.StatusUnauthorized, responseError{Message: err.Error()})
		// 	return
		// }
		authorizationHeaderStr := string(authorizationHeader)
		if len(authorizationHeaderStr) == 0 {
			err := errors.New("authorization header is not provided")
			c.AbortWithStatusJSON(http.StatusUnauthorized, responseError{Message: err.Error()})
			return
		}

		fields := strings.Fields(authorizationHeaderStr)
		if len(fields) < 2 {
			err := errors.New("invalid authorization header format")
			c.AbortWithStatusJSON(http.StatusUnauthorized, responseError{Message: err.Error()})
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != AuthorizationTypeBearer {
			err := fmt.Errorf("unsupported authorization type %s", authorizationType)
			c.AbortWithStatusJSON(http.StatusUnauthorized, responseError{Message: err.Error()})
			return
		}

		accessToken := fields[1]
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, responseError{Message: err.Error()})
			return
		}
		c.Set(AuthorizationPayloadKey, payload)
		c.Next(ctx)
	}
}
