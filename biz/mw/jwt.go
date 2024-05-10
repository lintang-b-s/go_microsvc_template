package mw

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	jalvaJWt "github.com/dgrijalva/jwt-go"
	gojwt "github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"

	utilJWT "lintang/go_hertz_template/biz/util/jwt"

	"github.com/hertz-contrib/jwt"
	"go.uber.org/zap"
)

var (
	// JwtMiddleware       *jwt.HertzJWTMiddleware
	IdentityKey = "id"
)

func GetJwtMiddleware() *jwt.HertzJWTMiddleware {

	pubKey, err := os.ReadFile("cert/id_rsa.pub")
	if err != nil {
		zap.L().Fatal("", zap.Error(err))
	}
	key, err := jalvaJWt.ParseRSAPublicKeyFromPEM(pubKey)
	if err != nil {
		zap.L().Fatal("", zap.Error(err))

	}

	JwtMiddleware, err := jwt.New(&jwt.HertzJWTMiddleware{
		Realm:       "dogker digital signature public key auth",
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: IdentityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*utilJWT.Payload); ok {
				return jwt.MapClaims{
					IdentityKey: v.ID,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			fmt.Println("IdentityHandler")

			claims := jwt.ExtractClaims(ctx, c)
			fmt.Println(claims)
			return &utilJWT.Payload{
				ID: claims[IdentityKey].(uuid.UUID),
			}
		},
		Authorizator: func(data interface{}, ctx context.Context, c *app.RequestContext) bool {
			fmt.Println("tes Authorizator")

			if v, ok := data.(*utilJWT.Payload); ok {
				fmt.Println(v)
				c.Set("userID", v.ID)

				return true
			}

			return false
		},
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			c.JSON(code, map[string]interface{}{
				"code":    code,
				"message": message,
			})
		},

		TokenLookup: "header: Authorization, query: token, cookie: jwt",

		// TokenHeadName is a string in the header. Default value is "Bearer". If you want empty value, use WithoutDefaultTokenHeadName.
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
		KeyFunc: func(t *gojwt.Token) (interface{}, error) {
			// if gojwt.GetSigningMethod(gojwt.SigningMethodECDSA) != t.Method {
			// 	return nil, jwt.ErrInvalidSigningAlgorithm
			// }
			fmt.Println("tess")
			if _, ok := t.Method.(*gojwt.SigningMethodRSA); !ok {
				return nil, jwt.ErrInvalidSigningAlgorithm
			}

			return key, nil
		},
	})
	if err != nil {
		zap.L().Fatal("JWT Error:"+err.Error(), zap.Error(err))
	}
	return JwtMiddleware
}
