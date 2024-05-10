package jwt

import (
	"errors"
	"lintang/go_hertz_template/config"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	gojwt "github.com/golang-jwt/jwt/v4"

	"go.uber.org/zap"
)

const minprivateKeySize = 32

// JWTMaker JWT maker
type JWTMaker struct {
	privateKey []byte
	publicKey  []byte
}

func NewJWTMaker(cfg *config.Config) *JWTMaker {
	prvKey, err := os.ReadFile("cert/id_rsa")
	if err != nil {
		zap.L().Fatal("", zap.Error(err))
	}
	pubKey, err := os.ReadFile("cert/id_rsa.pub")
	if err != nil {
		zap.L().Fatal("", zap.Error(err))
	}

	// if len(cfg.Auth.PrivateKey) < minprivateKeySize {
	// 	zap.L().Fatal(" len(cfg.Auth.PrivateKey) < minprivateKeySize ")
	// }
	return &JWTMaker{prvKey, pubKey}
}

// CreateToken membuat jwt token baru berdurasi utk user
func (maker *JWTMaker) CreateToken(userID string, username string, duration time.Duration) (string, *Payload, error) {

	key, err := jwt.ParseRSAPrivateKeyFromPEM(maker.privateKey)
	if err != nil {
		return "", nil, err
	}

	payload, err := NewPayload(userID,username, duration)
	if err != nil {
		return "", payload, err
	}

	jwtToken := gojwt.NewWithClaims(jwt.SigningMethodRS256, payload)
	token, err := jwtToken.SignedString(key)
	return token, payload, err
}

// VerifyToken cek jika token valid atau tidak
func (maker *JWTMaker) VerifyToken(token string) (*Payload, error) {

	key, err := jwt.ParseRSAPublicKeyFromPEM(maker.publicKey)
	if err != nil {
		return nil, err
	}
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodRSA)
		if !ok {
			return nil, ErrInvalidToken
		}
		return key, nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, ErrExpiredToken) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, ErrInvalidToken
	}

	return payload, nil
}
