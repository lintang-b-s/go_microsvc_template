package domain

import "time"

type Order struct {
	ID uint64 `json:"id"`
	UserID uint64 `json:"user_id"`
	PaymentID uint64 `json:"payment_id"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	DealerID uint64 `json:"dealer_id"`
	Price uint64 `json:"price"`
	NamaPembeli string `json:"nama_pembeli"`
	NomorTeleponPembeli string `json:"nomor_telepon_pembeli"`
	EmailPembeli string `json:"emailPembeli"`
	AlamatPembeli string `json:"alamat_pembeli"`
}


