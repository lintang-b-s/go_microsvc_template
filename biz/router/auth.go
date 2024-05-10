package router

import (
	"context"
	"lintang/go_hertz_template/biz/domain"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
)

type AuthService interface {
	Login(ctx context.Context, l domain.User) (domain.Token, error)
	RenewAccessToken(ctx context.Context, r NewTokenReq) (domain.NewToken, error)
	DeleteRefreshToken(ctx context.Context, d DeleteRefreshTokenReq) error
}

type AuthHandler struct {
	svc AuthService
}

func AuthRouter(r *server.Hertz, a AuthService) {
	handler := &AuthHandler{
		svc: a,
	}

	root := r.Group("/api/v1")
	{
		aH := root.Group("/auth")
		{
			aH.POST("/login", handler.login)
			aH.PUT("/token")
			aH.DELETE("/logout")
		}
	}

}

// auth

type LoginReq struct {
	Email    string `json:"email,required" vd:"email($)"`
	Password string `json:"password,required" vd:"password($)" `
}

func (h *AuthHandler) login(ctx context.Context, c *app.RequestContext) {
	var req LoginReq
	err := c.BindAndValidate(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseError{Message: err.Error()})
		return
	}
	res, err := h.svc.Login(ctx, domain.User{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, domain.Token{
		SessionId:             res.SessionId,
		AccessToken:           res.AccessToken,
		AccessTokenExpiresAt:  res.AccessTokenExpiresAt,
		RefreshToken:          res.RefreshToken,
		RefreshTokenExpiresAt: res.RefreshTokenExpiresAt,
		User: domain.User{
			ID:          res.User.ID,
			Username:    res.User.Username,
			Email:       res.User.Email,
			Dob:         res.User.Dob,
			Gender:      res.User.Gender,
			CreatedTime: res.User.CreatedTime,
		},
	})
}

type NewTokenReq struct {
	RefreshToken string `json:"refresh_token,required" vd:" regexp('^[\\w\\-\\.]*$')"`
}

func (h *AuthHandler) RenewAcctoken(ctx context.Context, c *app.RequestContext) {
	var req NewTokenReq
	err := c.BindAndValidate(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseError{Message: err.Error()})
		return
	}
	res, err := h.svc.RenewAccessToken(ctx, NewTokenReq{RefreshToken: req.RefreshToken})
	if err != nil {
		c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, domain.NewToken{AccessToken: res.AccessToken, AccessTokenExpiresAt: res.AccessTokenExpiresAt})
}

type DeleteRefreshTokenReq struct {
	RefreshToken string `json:"refresh_token,required" vd:" regexp('^[\\w\\-\\.]*$')"`
}

func (h *AuthHandler) Logout(ctx context.Context, c *app.RequestContext) {
	var req DeleteRefreshTokenReq
	err := c.BindAndValidate(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseError{Message: err.Error()})
		return
	}
	err = h.svc.DeleteRefreshToken(ctx, req)
	if err != nil {
		c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, "ok")
}
