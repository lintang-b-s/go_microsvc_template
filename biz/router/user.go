package router

import (
	"context"
	"errors"
	"lintang/go_hertz_template/biz/domain"
	"lintang/go_hertz_template/biz/mw"
	"lintang/go_hertz_template/biz/util/jwt"
	"net/http"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"go.uber.org/zap"
)

type UserService interface {
	Create(ctx context.Context, d domain.User) error
	Get(ctx context.Context, userID int32) (domain.User, error)
}

type UserHandler struct {
	svc UserService
	jwt jwt.JwtTokenMaker
}

func UserRouter(r *server.Hertz, u UserService, jwt jwt.JwtTokenMaker) {
	handler := &UserHandler{
		svc: u,
		jwt: jwt,
	}

	root := r.Group("/api/v1")
	{
		uH := root.Group("/user")
		{
			uH.POST("/", handler.Create)
			uH.GET("/", mw.Protected(jwt), handler.Get)
		}

	}
}

// ResponseError represent the response error struct
type ResponseError struct {
	Message string `json:"message"`
}

type createUserReq struct {
	Username string        `json:"username,required" vd:"len($) < 25 && regexp('^[\\w\\-\\.]*$'); msg:'nama harus alpahnumeric atau boleh juga mendandung simbol -,.'"`
	Password string        `json:"password,required" vd:" password($); msg:'password harus terdiri dari minimal 1 uppercase, 1 symbol , dan satu digit angka, dan panjangnnya antara 8-16'"`
	Email    string        `json:"email,required" vd:"email($)"`
	Gender   domain.Gender `json:"gender,required" vd:"in($, 'Male', 'Female'); msg:'jenis kelamin harus MALE atau FEMALE'"`
	Age      uint64        `json:"age,required" vd:" $ >= 10; msg:'umur harus lebih dari 10'"`
	Address  string        `json:"address,required"`
}

type createUserRes struct {
	Message string `json:"string"`
}

func (h *UserHandler) Create(ctx context.Context, c *app.RequestContext) {
	var req createUserReq
	err := c.BindAndValidate(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseError{Message: err.Error()})
		return
	}
	err = h.svc.Create(ctx, domain.User{
		Username: req.Username,
		Email:    req.Email,
		Gender:   req.Gender,
		Age:      req.Age,
		Address:  req.Address,
	})
	if err != nil {
		c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, createUserRes{Message: "user creatd successfully"})
}

type userRes struct {
	User domain.User `json:"user"`
}

func (h *UserHandler) Get(ctx context.Context, c *app.RequestContext) {
	authPayload := c.MustGet(mw.AuthorizationPayloadKey).(*jwt.Payload)
	userID := authPayload.ID
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		zap.L().Error("strconv.Atoi (Get) (UserHandler)", zap.Error(err))
		c.JSON(consts.StatusInternalServerError, ResponseError{Message: err.Error()})
	}
	user, err := h.svc.Get(ctx, int32(userIDInt))
	if err != nil {

		c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, userRes{user})
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}
	var ierr *domain.Error
	if !errors.As(err, &ierr) {
		return http.StatusInternalServerError
	} else {
		switch ierr.Code() {
		case domain.ErrInternalServerError:
			return http.StatusInternalServerError
		case domain.ErrNotFound:
			return http.StatusNotFound
		case domain.ErrConflict:
			return http.StatusConflict
		case domain.ErrBadParamInput:
			return http.StatusBadRequest
		default:
			return http.StatusInternalServerError
		}
	}

}
