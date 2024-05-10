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

type SessionRepository struct {
	db *Postgres
}

func NewSessionRepo(db *Postgres) *SessionRepository {
	return &SessionRepository{db}
}

func (r *SessionRepository) Insert(ctx context.Context, s domain.Session) error {
	q := queries.New(r.db.Pool)

	err := q.InsertSession(ctx, queries.InsertSessionParams{
		ID:           googleuuid.UUID(s.ID),
		Username:     s.Username,
		RefreshToken: s.RefreshToken,
		ExpiresAt:    pgtype.Timestamptz{Valid: true, Time: s.ExpiresAt},
	})
	if err != nil {
		zap.L().Error("InsertSession (SessionRepository)", zap.Error(err))
		return domain.WrapErrorf(err, domain.ErrInternalServerError, domain.MessageInternalServerError)
	}

	return nil
}

func (r *SessionRepository) Get(ctx context.Context, reftokenID string) (domain.Session, error) {
	sessionUUID, err := uuid.FromString(reftokenID)
	if err != nil {
		zap.L().Error("uuid.FromString (SessionRepository)", zap.Error(err))
		return domain.Session{}, domain.WrapErrorf(err, domain.ErrInternalServerError, domain.MessageInternalServerError)
	}

	q := queries.New(r.db.Pool)
	s, err := q.GetSession(ctx, googleuuid.UUID(sessionUUID))
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
		ExpiresAt:    s.ExpiresAt.Time,
	}
	return res, nil
}

func (r *SessionRepository) Delete(ctx context.Context, sID string) error {
	sessionUUID, err := uuid.FromString(sID)
	if err != nil {
		zap.L().Error("uuid.FromString (SessionRepository)", zap.Error(err))
		return domain.WrapErrorf(err, domain.ErrInternalServerError, domain.MessageInternalServerError)
	}
	q := queries.New(r.db.Pool)

	err = q.DeleteSession(ctx, googleuuid.UUID(sessionUUID))
	if err != nil {
		return domain.WrapErrorf(err, domain.ErrInternalServerError, domain.MessageInternalServerError)
	}
	return nil
}
