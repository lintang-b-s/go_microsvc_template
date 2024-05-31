package db

import (
	"context"
	"lintang/go_hertz_template/biz/dal/db/queries"
	"lintang/go_hertz_template/biz/domain"

	"go.uber.org/zap"
)

type CarAdsRepository struct {
	db *Mysql
}

func NewCarAdsRepo(db *Mysql) *CarAdsRepository {
	return &CarAdsRepository{db}
}

func (r *CarAdsRepository) GetCarAds(ctx context.Context) ([]domain.CarAdsHomePage, error) {
	q := queries.New(r.db.Conn)

	ads, err := q.HomePage(ctx)
	if err != nil {
		zap.L().Error("q.HomePage (GetCarAds) (CarAdsRepository)", zap.Error(err))
		return []domain.CarAdsHomePage{}, err
	}

	var homePageAds []domain.CarAdsHomePage
	for i := 0; i < len(ads); i++ {
		homePageAds = append(homePageAds, domain.CarAdsHomePage{
			ID:             uint64(ads[i].ID),
			CarID:          uint64(ads[i].CarID),
			DealerID:       uint64(ads[i].DealerID),
			Title:          ads[i].Title,
			Description:    ads[i].Description,
			IsActive:       ads[i].IsActive,
			PostingDate:    ads[i].PostingDate,
			YearProduction: uint64(ads[i].YearProduction),
			Color:          ads[i].Color,
			Mileage:        uint64(ads[i].Mileage),
			Price:          uint64(ads[i].Price),
			Status:         domain.CarAdsStatus(ads[i].Status.CarAdvertisementsStatus),
			Location:       ads[i].Location,
			BahanBakar:     domain.BahanBakar(ads[i].BahanBakar),
			TransType:      domain.TransType(ads[i].TransType),
			KapasitasMesin: domain.KapasitasMesin(ads[i].KapasitasMesin),
			Stok:           uint64(ads[i].Stok),
			Brand:          ads[i].Brand,
			Model:          ads[i].Model,
			CarType:        domain.CarType(ads[i].CarType),
			ThumbnailImage: ads[i].Thumbnailimage.String,
		})
	}

	return homePageAds, nil
}

func (r *CarAdsRepository) ShowAdsByID(ctx context.Context, adsID uint64) (domain.CarAd, error) {
	q := queries.New(r.db.Conn)

	ad, err := q.ShowAdsByID(ctx, int32(adsID))
	if err != nil {
		zap.L().Error("q.ShowAdsByID (ShowAdsByID) (CarAdsRepository)", zap.Error(err))
		return domain.CarAd{}, err
	}
	carAd := domain.CarAd{
		ID:             uint64(ad.ID),
		CarID:          uint64(ad.CarID),
		DealerID:       uint64(ad.DealerID),
		Title:          ad.Title,
		Description:    ad.Description,
		IsActive:       ad.IsActive,
		PostingDate:    ad.PostingDate,
		YearProduction: uint64(ad.YearProduction),
		Color:          ad.Color,
		Mileage:        uint64(ad.Mileage),
		Price:          uint64(ad.Price),
		Status:         domain.CarAdsStatus(ad.Status.CarAdvertisementsStatus),
		Location:       ad.Location,
		BahanBakar:     domain.BahanBakar(ad.BahanBakar),
		TransType:      domain.TransType(ad.TransType),
		KapasitasMesin: domain.KapasitasMesin(ad.KapasitasMesin),
		Stok:           uint64(ad.Stok),
		Brand:          ad.Brand,
		Model:          ad.Model,
		CarType:        domain.CarType(ad.CarType),
		ThumbnailImage: ad.Thumbnailimage.String,
		DealerName:     ad.DealerName,
		NoTelp:         ad.NoTelp,
		Email:          ad.Email,
		LocationDealer: ad.Location,
	}
	return carAd, nil
}

func (r *CarAdsRepository) FilterCar(ctx context.Context, brand string, color string, location string,
	bahanBakar domain.BahanBakar, transType domain.TransType, fromYearProduction int32, toYearProduction int32,
	fromPrice int32, toPrice int32, fromMileage int32, toMileage int32) ([]domain.FilteredAd, error) {
	q := queries.New(r.db.Conn)
	adsFilteredRow, err := q.FilterCar(ctx, queries.FilterCarParams{
		Column1:            brand,
		Brand:              brand,
		Color:              color,
		Column4:            color,
		Location:           location,
		Column6:            location,
		BahanBakar:         queries.CarAdvertisementsBahanBakar(bahanBakar),
		Column8:            queries.CarAdvertisementsBahanBakar(bahanBakar),
		TransType:          queries.CarAdvertisementsTransType(transType),
		Column10:           queries.CarAdvertisementsTransType(transType),
		FromYearProduction: fromYearProduction,
		ToYearProduction:   toYearProduction,
		FromPrice:          fromPrice,
		ToPrice:            toPrice,
		FromMileage:        fromMileage,
		ToMileage:          toMileage,
	})
	if err != nil {
		zap.L().Error("q.FilterCar (FilterCar) (CarAdsRepository)", zap.Error(err))
		return []domain.FilteredAd{}, err
	}

	var filteredAds []domain.FilteredAd

	for i := 0; i < len(adsFilteredRow); i++ {
		filteredAds = append(filteredAds, domain.FilteredAd{
			ID:             adsFilteredRow[i].ID,
			Brand:          adsFilteredRow[i].Brand,
			Model:          adsFilteredRow[i].Model,
			YearProduction: adsFilteredRow[i].YearProduction,
			Price:          adsFilteredRow[i].Price,
			Mileage:        adsFilteredRow[i].Mileage,
			Location:       adsFilteredRow[i].Location,
			BahanBakar:     domain.BahanBakar(adsFilteredRow[i].BahanBakar),
			TransType:      domain.TransType(adsFilteredRow[i].TransType),
		})
	}

	return filteredAds, nil
}

//  buat compare car

func (r *CarAdsRepository) GetActiveAds(ctx context.Context) ([]string, error) {
	q := queries.New(r.db.Conn)
	carTitles, err := q.GetActiveAds(ctx)
	if err != nil {
		zap.L().Error("q.GetActiveAds (GetActiveAds) (CarAdsRepository)", zap.Error(err ))
		return []string{}, err 
	}
	return carTitles, nil 
}

func (r *CarAdsRepository) CompareCar(ctx context.Context, titleOne, titleTwo string) ([]domain.CarAds, error) {
	q := queries.New(r.db.Conn)
	carOne, err := q.GetAdsByTitle(ctx, titleOne)
	if err != nil {
		zap.L().Error("q.GetAdsByTitle (CompareCar) (CarAdsRepo)", zap.Error(err))
		return []domain.CarAds{}, err
	}

	carTwo, err := q.GetAdsByTitle(ctx, titleTwo)
	if err != nil {
		zap.L().Error("q.GetAdsByTitle (CompareCar) (CarAdsRepo)", zap.Error(err))
		return []domain.CarAds{}, err
	}

	var twoCars []domain.CarAds
	twoCars = append(twoCars, domain.CarAds{
		ID:             uint64(carOne.ID),
		CarID:          uint64(carOne.CarID),
		DealerID:       uint64(carOne.DealerID),
		Title:          carOne.Title,
		Description:    carOne.Description,
		IsActive:       carOne.IsActive,
		PostingDate:    carOne.PostingDate,
		YearProduction: uint64(carOne.YearProduction),
		Color:          carOne.Color,
		Mileage:        uint64(carOne.Mileage),
		Price:          uint64(carOne.Price),
		Status:         domain.CarAdsStatus(carOne.Status.CarAdvertisementsStatus),
		Location:       carOne.Location,
		BahanBakar:     domain.BahanBakar(carOne.BahanBakar),
		TransType:      domain.TransType(carOne.TransType),
		KapasitasMesin: domain.KapasitasMesin(carOne.KapasitasMesin),
		Stok:           uint64(carOne.Stok),

		ThumbnailImage: carOne.Thumbnailimage.String,
	})

	twoCars = append(twoCars, domain.CarAds{
		ID:             uint64(carTwo.ID),
		CarID:          uint64(carTwo.CarID),
		DealerID:       uint64(carTwo.DealerID),
		Title:          carTwo.Title,
		Description:    carTwo.Description,
		IsActive:       carTwo.IsActive,
		PostingDate:    carTwo.PostingDate,
		YearProduction: uint64(carTwo.YearProduction),
		Color:          carTwo.Color,
		Mileage:        uint64(carTwo.Mileage),
		Price:          uint64(carTwo.Price),
		Status:         domain.CarAdsStatus(carTwo.Status.CarAdvertisementsStatus),
		Location:       carTwo.Location,
		BahanBakar:     domain.BahanBakar(carTwo.BahanBakar),
		TransType:      domain.TransType(carTwo.TransType),
		KapasitasMesin: domain.KapasitasMesin(carTwo.KapasitasMesin),
		Stok:           uint64(carTwo.Stok),

		ThumbnailImage: carTwo.Thumbnailimage.String,
	})

	return twoCars, nil
}


