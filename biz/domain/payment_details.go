package domain

import "time"

type PaymentMethod string
const (
	Paypal PaymentMethod =  "Paypal"
	CreditCard PaymentMethod = "Credit Card"
	Cash PaymentMethod = "Cash"
	Loan PaymentMethod = "Loan"
	DebitCard PaymentMethod = "Debit Card"
)


type PaymentStatus string
const (
	PAID PaymentStatus = "Paid"
	PENDING PaymentStatus = "Pending"
	CANCELED PaymentStatus = "Canceled"
	EXPIRED PaymentStatus = "Expired"
)
type PaymentDetail struct {
	ID uint64 `json:"id"`
	Amount uint64 `json:"amount"`
	PaymentMethod  PaymentMethod `json:"payment_method"`
	Status PaymentStatus `json:"status"`
	Provider string `json:"provider"`
	CreatedAt time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
}


