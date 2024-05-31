package domain

type Dealer struct {
	ID uint64 `json:"id"`
	DealerName string `json:"dealer_name"`
	NoTelp string `json:"no_telp"`
	Email string `json:"email"`
	Location string `json:"location"`
	Password string `json:"password"`
}


