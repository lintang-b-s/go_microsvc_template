package db

import (
	"context"
	"fmt"
	"lintang/go_hertz_template/biz/dal/db/queries"
	"lintang/go_hertz_template/biz/domain"

	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

type SessionRepository struct {
	db *Mysql
}

func NewSessionRepo(db *Mysql) *SessionRepository {
	return &SessionRepository{db}
}

func (r *SessionRepository) Insert(ctx context.Context, s domain.Session) error {
	q := queries.New(r.db.Conn)

	err := q.InsertSession(ctx, queries.InsertSessionParams{
		RefTokenID:   s.ID,
		Username:     s.Username,
		RefreshToken: s.RefreshToken,
		ExpiresAt:    s.ExpiresAt,
		// pgtype.Timestamptz{Valid: true, Time: s.ExpiresAt},
	})
	if err != nil {
		zap.L().Error("q.InsertSession (SessionRepository)", zap.Error(err))
		return domain.WrapErrorf(err, domain.ErrInternalServerError, domain.MessageInternalServerError)
	}

	return nil
}

func (r *SessionRepository) Get(ctx context.Context, reftokenID string) (domain.Session, error) {

	q := queries.New(r.db.Conn)
	s, err := q.GetSession(ctx, reftokenID)
	if err != nil {
		if err == pgx.ErrNoRows {
			zap.L().Debug(fmt.Sprint("session with id  %s not exists", reftokenID))
			return domain.Session{}, domain.WrapErrorf(err, domain.ErrNotFound, fmt.Sprintf("session with id  %s not exists", reftokenID))
		}
		zap.L().Error("q.GetSession (SessionRepositroy)", zap.Error(err))
		return domain.Session{}, domain.WrapErrorf(err, domain.ErrInternalServerError, domain.MessageInternalServerError)
	}
	res := domain.Session{
		ID:           s.ID,
		Username:     s.Username,
		RefreshToken: s.RefreshToken,
		ExpiresAt:    s.ExpiresAt,
	}
	return res, nil
}

func (r *SessionRepository) Delete(ctx context.Context, sID string) error {

	q := queries.New(r.db.Conn)

	err := q.DeleteSession(ctx, sID)
	if err != nil {
		return domain.WrapErrorf(err, domain.ErrInternalServerError, domain.MessageInternalServerError)
	}
	return nil
}
