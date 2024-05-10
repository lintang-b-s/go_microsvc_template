package db

import (
	"context"
	"fmt"
	"lintang/go_hertz_template/biz/dal/db/queries"
	"lintang/go_hertz_template/biz/domain"

	"github.com/gofrs/uuid"
	googleuuid "github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"go.uber.org/zap"
)

type UserRepository struct {
	db *Postgres
}

func NewUserRepo(db *Postgres) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) Insert(ctx context.Context, u domain.User) error {
	q := queries.New(r.db.Pool)

	var gender queries.Gender
	gender = queries.GenderFemale
	if u.Gender == domain.Male {
		gender = queries.GenderMale
	}

	_, err := q.InsertUser(ctx, queries.InsertUserParams{
		Username: u.Username,
		Email:    u.Email,
		Dob:      pgtype.Date{Valid: true, Time: u.Dob},
		Gender:   gender,
		Password: u.Password,
	})
	if err != nil {
		zap.L().Error("InsertUser (UserRepository)", zap.Error(err))
		return domain.WrapErrorf(err, domain.ErrInternalServerError, domain.MessageInternalServerError)
	}
	return nil
}

func (r *UserRepository) Get(ctx context.Context, userID string) (domain.User, error) {
	q := queries.New(r.db.Pool)
	userIDUUID, err := uuid.FromString(userID)
	if err != nil {
		zap.L().Error("uuid.FromString (UserRepository)", zap.Error(err))
		return domain.User{}, domain.WrapErrorf(err, domain.ErrInternalServerError, domain.MessageInternalServerError)
	}

	u, err := q.GetUser(ctx, googleuuid.UUID(userIDUUID))
	if err != nil {
		if err == pgx.ErrNoRows {
			zap.L().Debug(fmt.Sprint("user with id  %s not exists", userID))
			return domain.User{}, domain.WrapErrorf(err, domain.ErrNotFound, fmt.Sprintf("user with id  %s not exists", userID))
		}
		zap.L().Error("q.GetUser (UserRepistory)", zap.Error(err))
		return domain.User{}, domain.WrapErrorf(err, domain.ErrInternalServerError, domain.MessageInternalServerError)
	}

	d := domain.User{
		ID:          u.ID.String(),
		Username:    u.Username,
		Email:       u.Email,
		Dob:         u.Dob.Time,
		Gender:      domain.Gender(u.Gender),
		CreatedTime: u.CreatedTime.Time,
		UpdatedTime: u.UpdatedTime.Time,
	}
	return d, nil
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (domain.User, error) {

	q := queries.New(r.db.Pool)

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
		ID: u.ID.String(),
		Username:    u.Username,
		Email:       u.Email,
		Dob:         u.Dob.Time,
		Gender:      domain.Gender(u.Gender),
		CreatedTime: u.CreatedTime.Time,
		UpdatedTime: u.UpdatedTime.Time,
		Password: u.Password,
	}
	return d, nil
}
