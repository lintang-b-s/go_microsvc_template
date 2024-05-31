package service

import (
	"context"
	"lintang/go_hertz_template/biz/domain"

	"go.uber.org/zap"
)

type OrderRepo interface {
	CreateOrder(ctx context.Context, userID uint64,
		adsID uint64, paymentMethod domain.PaymentMethod, provider string,
		quantity uint64, price uint64, alamat string,
		telepon string, email string, nama string) error
}

type OrderService struct {
	orderRepo OrderRepo
}

func NewOrderService(o OrderRepo) *OrderService {
	return &OrderService{o}
}

func (c *OrderService) InsertOrder(ctx context.Context, userID uint64,
	adsID uint64, paymentMethod domain.PaymentMethod, provider string,
	quantity uint64, price uint64, alamat string,
	telepon string, email string, nama string) error {

	err := c.orderRepo.CreateOrder(ctx, userID,
		adsID, paymentMethod, provider,
		quantity, price, alamat,
		telepon, email, nama)

	if err != nil {
		zap.L().Error(" c.orderRepo.CreateOrder (InsertOrder) (OrderRepo)", zap.Error(err))
		return err
	}

	return nil
}
