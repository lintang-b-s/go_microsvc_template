package router

import (
	"context"
	"lintang/go_hertz_template/biz/domain"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
)

type AdsService interface {
	GetHomePageAds(ctx context.Context) ([]domain.CarAdsHomePage, error)
	ShowAdsByID(ctx context.Context, adsID uint64) (domain.CarAd, error)
	FilterCar(ctx context.Context, req FilterCarReq) ([]domain.FilteredAd, error)
	GetActiveAds(ctx context.Context) ([]string, error)
	CompareCar(ctx context.Context, titleOne, titleTwo string) ([]domain.CarAds, error)
}

type AdsHandler struct {
	svc AdsService
}

func AdsRouter(r *server.Hertz, a AdsService) {
	handler := &AdsHandler{
		svc: a,
	}

	root := r.Group("/api/v1")
	{
		uH := root.Group("/ads")
		{
			uH.GET("/", handler.GetHomePageAds)
			// uH.GET("/", mw.Protected(jwt), handler.Get)
		}

	}

}

type homePageAdsRes struct {
	Ads []domain.CarAdsHomePage `json:"ads"`
}

func (h *AdsHandler) GetHomePageAds(ctx context.Context, c *app.RequestContext) {
	ads, err := h.svc.GetHomePageAds(ctx)
	if err != nil {
		c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, homePageAdsRes{ads})
}

type showAdsByIDReq struct {
	AdsID uint64 `json:"adsID"`
}

type showAdsByIDRes struct {
	Ads domain.CarAd `json:"ads"`
}

func (h *AdsHandler) ShowAdsByID(ctx context.Context, c *app.RequestContext) {
	var req showAdsByIDReq
	err := c.BindAndValidate(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseError{Message: err.Error()})
		return
	}
	ads, err := h.svc.ShowAdsByID(ctx, req.AdsID)
	if err != nil {
		c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, showAdsByIDRes{Ads: ads})
}

type FilterCarReq struct {
	Brand              string
	Color              string
	Location           string
	BahanBakar         domain.BahanBakar
	TransType          domain.TransType
	FromYearProduction int32
	ToYearProduction   int32
	FromPrice          int32
	ToPrice            int32
	FromMileage        int32
	ToMileage          int32
}

type filterCarRes struct {
	Ads []domain.FilteredAd `json:"filtered_ads"`
}

func (h *AdsHandler) FilterCar(ctx context.Context, c *app.RequestContext) {
	var req FilterCarReq
	err := c.BindAndValidate(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseError{Message: err.Error()})
		return
	}

	ads, err := h.svc.FilterCar(ctx, req)

	if err != nil {
		c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, filterCarRes{Ads: ads})
}
