package service

import (
	"context"
	"lintang/go_hertz_template/biz/domain"
	"lintang/go_hertz_template/biz/util"
)

type UserRepository interface {
	Insert(ctx context.Context, u domain.User) error
	Get(ctx context.Context, userID string) (domain.User, error)
	GetByEmail(ctx context.Context, email string) (domain.User, error)
}

type UserService struct {
	userRepo UserRepository
}

func NewUserService(u UserRepository) *UserService {
	return &UserService{
		userRepo: u,
	}
}

func (u *UserService) Create(ctx context.Context, d domain.User) error {
	d.Password, _ = util.HashPassword(d.Password)
	err := u.userRepo.Insert(ctx, d)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserService) Get(ctx context.Context, userID string) (domain.User, error) {
	user, err := u.userRepo.Get(ctx, userID)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}
