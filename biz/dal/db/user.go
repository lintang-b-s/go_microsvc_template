package db

import (
	"context"
	"fmt"
	"lintang/go_hertz_template/biz/dal/db/queries"
	"lintang/go_hertz_template/biz/domain"

	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

type UserRepository struct {
	db *Mysql
}

func NewUserRepo(db *Mysql) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) Insert(ctx context.Context, u domain.User) error {
	q := queries.New(r.db.Conn)

	var gender queries.UsersGender
	gender = queries.UsersGenderFemale
	if u.Gender == domain.Male {
		gender = queries.UsersGenderMale
	}

	err := q.InsertUser(ctx, queries.InsertUserParams{
		UserName: u.Username,
		Email:    u.Email,
		Gender:   gender,
		Password: u.Password,
		Age:      int32(u.Age),
		Address:  u.Address,
	})
	if err != nil {
		zap.L().Error("InsertUser (UserRepository)", zap.Error(err))
		return domain.WrapErrorf(err, domain.ErrInternalServerError, domain.MessageInternalServerError)
	}
	return nil
}

func (r *UserRepository) Get(ctx context.Context, userID int32) (domain.User, error) {
	q := queries.New(r.db.Conn)

	u, err := q.GetUser(ctx, userID)
	if err != nil {
		if err == pgx.ErrNoRows {
			zap.L().Debug(fmt.Sprint("user with id  %s not exists", userID))
			return domain.User{}, domain.WrapErrorf(err, domain.ErrNotFound, fmt.Sprintf("user with id  %s not exists", userID))
		}
		zap.L().Error("q.GetUser (UserRepistory)", zap.Error(err))
		return domain.User{}, domain.WrapErrorf(err, domain.ErrInternalServerError, domain.MessageInternalServerError)
	}

	d := domain.User{
		ID:       uint64(u.ID),
		Username: u.UserName,
		Email:    u.Email,
		Gender:   domain.Gender(u.Gender),
		Age:      uint64(u.Age),
		Address:  u.Address,
	}
	return d, nil
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (domain.User, error) {

	q := queries.New(r.db.Conn)

	u, err := q.GetUserByEmail(ctx, email)
	if err != nil {
		if err == pgx.ErrNoRows {
			zap.L().Debug(fmt.Sprint("user with email  %s not exists", email))
			return domain.User{}, domain.WrapErrorf(err, domain.ErrNotFound, fmt.Sprintf("user with email %s not exists", email))
		}
		zap.L().Error("q.GetUser (UserRepistory)", zap.Error(err))
		return domain.User{}, domain.WrapErrorf(err, domain.ErrInternalServerError, domain.MessageInternalServerError)
	}

	d := domain.User{
		ID:       uint64(u.ID),
		Username: u.UserName,
		Email:    u.Email,
		Gender:   domain.Gender(u.Gender),
		Password: u.Password,
		Age:      uint64(u.Age),
		Address:  u.Address,
	}
	return d, nil
}
