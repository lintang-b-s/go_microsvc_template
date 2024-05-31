package db

import (
	"context"
	"database/sql"
	"lintang/go_hertz_template/biz/dal/db/queries"
	"lintang/go_hertz_template/biz/domain"

	"go.uber.org/zap"
)

type OrderRepo struct {
	db *Mysql
}

func NewOrderRepo(db *Mysql) *OrderRepo {
	return &OrderRepo{db}
}

func (r *CarAdsRepository) CreateOrder(ctx context.Context, userID uint64,
	adsID uint64, paymentMethod domain.PaymentMethod, provider string,
	quantity uint64, price uint64, alamat string,
	telepon string, email string, nama string) error {

	q := queries.New(r.db.Conn)

	ads, err := q.SelectAdsByID(ctx, int32(adsID))
	if err != nil {
		zap.L().Error("q.SelectAdsByID (CreateOrder) (OrderRepo)", zap.Error(err))
		return err
	}

	err = q.InsertPaymentDetails(ctx, queries.InsertPaymentDetailsParams{
		Amount:        ads.Price * int32(quantity),
		PaymentMethod: queries.NullPaymentDetailsPaymentMethod{Valid: true, PaymentDetailsPaymentMethod: queries.PaymentDetailsPaymentMethod(paymentMethod)},
		Status:        "Pending",
		Provider:      provider,
	})

	if err != nil {
		zap.L().Error("q.InsertPaymentDetails (CreateOrder) (OrderRepo)", zap.Error(err))
		return err
	}

	paymentID, err := q.GetInsertedPaymentID(ctx)
	if err != nil {
		zap.L().Error("q.GetInsertedPaymentID (CreateOrder) (OrderRepo)", zap.Error(err))
		return err
	}

	dealerID := ads.DealerID

	err = q.UpdateAds(ctx, queries.UpdateAdsParams{
		Stok: int32(quantity),
		ID:   ads.ID,
	})
	if err != nil {
		zap.L().Error("q.UpdateAds (CreateOrder) (OrderRepo)", zap.Error(err))
		return err
	}

	err = q.InsertOrder(ctx, queries.InsertOrderParams{
		UserID:              int32(userID),
		PaymentID:           int32(paymentID),
		Price:               sql.NullInt32{Valid: true, Int32: ads.Price * int32(quantity)},
		AlamatPembeli:       sql.NullString{Valid: true, String: alamat},
		NomorTeleponPembeli: sql.NullString{Valid: true, String: telepon},
		Emailpembeli:        sql.NullString{Valid: true, String: email},
		NamaPembeli:         sql.NullString{Valid: true, String: nama},
		DealerID:            sql.NullInt32{Valid: true, Int32: dealerID},
	})

	if err != nil {
		zap.L().Error("q.InsertOrder (CreateOrder) (OrderRepo)", zap.Error(err))
		return err
	}

	return nil
}


