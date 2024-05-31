package service

import (
	"context"
	"lintang/go_hertz_template/biz/domain"
	"lintang/go_hertz_template/biz/router"

	"go.uber.org/zap"
)

type CarAdsRepository interface {
	GetCarAds(ctx context.Context) ([]domain.CarAdsHomePage, error)
	ShowAdsByID(ctx context.Context, adsID uint64) (domain.CarAd, error)
	FilterCar(ctx context.Context, brand string, color string, location string,
		bahanBakar domain.BahanBakar, transType domain.TransType, fromYearProduction int32, toYearProduction int32,
		fromPrice int32, toPrice int32, fromMileage int32, toMileage int32) ([]domain.FilteredAd, error)
	GetActiveAds(ctx context.Context) ([]string, error)
	CompareCar(ctx context.Context, titleOne, titleTwo string) ([]domain.CarAds, error)
}

type CarAdsService struct {
	adsRepo CarAdsRepository
}

func NewCarAdsService(c CarAdsRepository) *CarAdsService {
	return &CarAdsService{c}
}

func (c *CarAdsService) GetHomePageAds(ctx context.Context) ([]domain.CarAdsHomePage, error) {
	ads, err := c.adsRepo.GetCarAds(ctx)
	if err != nil {
		zap.L().Error("c.adsRepo.GetCarAds (GetHomePageAds) (CarAdsServcice)", zap.Error(err))
		return []domain.CarAdsHomePage{}, err
	}
	return ads, nil
}

func (c *CarAdsService) ShowAdsByID(ctx context.Context, adsID uint64) (domain.CarAd, error) {
	ads, err := c.adsRepo.ShowAdsByID(ctx, adsID)
	if err != nil {

		return domain.CarAd{}, err
	}
	return ads, nil
}

func (c *CarAdsService) FilterCar(ctx context.Context, req router.FilterCarReq) ([]domain.FilteredAd, error) {

	filteredAds, err := c.adsRepo.FilterCar(ctx, req.Brand, req.Color, req.Location, req.BahanBakar, req.TransType, req.FromYearProduction,
		req.ToYearProduction, req.FromPrice, req.ToPrice, req.FromMileage, req.ToMileage)
	if err != nil {
		return []domain.FilteredAd{}, err
	}
	return filteredAds, nil
}

func (c *CarAdsService) GetActiveAds(ctx context.Context) ([]string, error) {
	carTitles, err := c.adsRepo.GetActiveAds(ctx)
	if err != nil {
		return []string{}, err 
	}
	return carTitles, nil
}

func  (c *CarAdsService) CompareCar(ctx context.Context, titleOne, titleTwo string) ([]domain.CarAds, error) {
	twoCar ,err := c.adsRepo.CompareCar(ctx, titleOne, titleTwo)
	if err != nil {
		return []domain.CarAds{}, err 
	}
	return twoCar, nil 
}






