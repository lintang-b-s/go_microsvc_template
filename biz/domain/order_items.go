package domain

import "time"

type OrderItems struct {
	ID uint64 `json:"id"`
	OrderID uint64 `json:"order_id"`
	AdvertisementID uint64 `json:"advertisement_id"`
	Quantity uint64 `json:"quantity"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	Price uint64 `json:"price"`
}



