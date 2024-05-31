package domain

import (
	"time"
)

type CarAdsStatus string

const (
	CarAdsNew    CarAdsStatus = "New"
	CarAdsSecond CarAdsStatus = "Second"
)

type BahanBakar string

const (
	BBBensin   BahanBakar = "Bensin"
	BBHybrid   BahanBakar = "Hybrid"
	BBElectric BahanBakar = "Electric"
)

type TransType string

const (
	TTAutomatic TransType = "Automatic"
	TTManual    TransType = "Manual"
)

type KapasitasMesin string

const (
	LowerOneThousand               KapasitasMesin = "<1000 cc"
	MoreThanOneThousand            KapasitasMesin = ">1000 - 1500 cc"
	MoreThanOneThousandFiveHundred KapasitasMesin = ">1500 - 2000 cc"
	MoreThanTwoThousand            KapasitasMesin = ">2000 - 3000 cc"
)

type CarAds struct {
	ID             uint64         `json:"id"`
	CarID          uint64         `json:"car_id"`
	DealerID       uint64         `json:"dealer_id"`
	Title          string         `json:"title"`
	Description    string         `json:"description"`
	IsActive       bool           `json:"is_active"`
	PostingDate    time.Time      `json:"posting_date"`
	YearProduction uint64         `json:"yaer_production"`
	Color          string         `json:"color"`
	Mileage        uint64         `json:"mileage"`
	Price          uint64         `json:"price"`
	Status         CarAdsStatus   `json:"status"`
	Location       string         `json:"location"`
	BahanBakar     BahanBakar     `json:"bahan_bakar"`
	TransType      TransType      `json:"trans_type"`
	KapasitasMesin KapasitasMesin `json:"kapasitas_mesin"`
	Stok           uint64         `json:"stok"`
	ThumbnailImage string         `json:"thumbnailImage"`
}

type CarAdsHomePage struct {
	ID             uint64         `json:"id"`
	CarID          uint64         `json:"car_id"`
	DealerID       uint64         `json:"dealer_id"`
	Title          string         `json:"title"`
	Description    string         `json:"description"`
	IsActive       bool           `json:"is_active"`
	PostingDate    time.Time      `json:"posting_date"`
	YearProduction uint64         `json:"yaer_production"`
	Color          string         `json:"color"`
	Mileage        uint64         `json:"mileage"`
	Price          uint64         `json:"price"`
	Status         CarAdsStatus   `json:"status"`
	Location       string         `json:"location"`
	BahanBakar     BahanBakar     `json:"bahan_bakar"`
	TransType      TransType      `json:"trans_type"`
	KapasitasMesin KapasitasMesin `json:"kapasitas_mesin"`
	Stok           uint64         `json:"stok"`
	Brand          string         `json:"brand"`
	Model          string         `json:"model"`
	CarType        CarType        `json:"car_type"`
	ThumbnailImage string         `json:"thumbnailImage"`
}

type CarAd struct {
	ID             uint64         `json:"id"`
	CarID          uint64         `json:"car_id"`
	DealerID       uint64         `json:"dealer_id"`
	Title          string         `json:"title"`
	Description    string         `json:"description"`
	IsActive       bool           `json:"is_active"`
	PostingDate    time.Time      `json:"posting_date"`
	YearProduction uint64         `json:"yaer_production"`
	Color          string         `json:"color"`
	Mileage        uint64         `json:"mileage"`
	Price          uint64         `json:"price"`
	Status         CarAdsStatus   `json:"status"`
	Location       string         `json:"location"`
	BahanBakar     BahanBakar     `json:"bahan_bakar"`
	TransType      TransType      `json:"trans_type"`
	KapasitasMesin KapasitasMesin `json:"kapasitas_mesin"`
	Stok           uint64         `json:"stok"`
	Brand          string         `json:"brand"`
	Model          string         `json:"model"`
	CarType        CarType        `json:"car_type"`
	ThumbnailImage string         `json:"thumbnailImage"`
	DealerName     string         `json:"dealer_name"`
	NoTelp         string         `json:"no_telp_dealer"`
	Email          string         `json:"email"`
	LocationDealer string         `json:"dealer_location"`
}

type FilteredAd struct {
	ID             int32
	Brand          string
	Model          string
	YearProduction int32
	Price          int32
	Mileage        int32
	Location       string
	BahanBakar     BahanBakar
	TransType      TransType
}
