package service

import (
	"context"
	"lintang/go_hertz_template/biz/domain"
	"lintang/go_hertz_template/biz/router"
	"lintang/go_hertz_template/biz/util"
	"lintang/go_hertz_template/biz/util/jwt"
	"time"
)

type SessionRepo interface {
	Insert(ctx context.Context, s domain.Session) error
	Get(ctx context.Context, reftokenID string) (domain.Session, error)
	Delete(ctx context.Context, sID string) error
}

type AuthService struct {
	jwtTokenMaker jwt.JwtTokenMaker
	sessionRepo   SessionRepo
	userRepo      UserRepository
}

func NewAuthService(sr SessionRepo, j jwt.JwtTokenMaker, ur UserRepository) *AuthService {
	return &AuthService{
		j, sr, ur,
	}
}

// Login: logic login use case
func (uc *AuthService) Login(ctx context.Context, l domain.User) (domain.Token, error) {
	user, err := uc.userRepo.GetByEmail(ctx, l.Email)
	if err != nil {
		// Bad request User with email not found in DB
		return domain.Token{}, err
	}

	err = util.CheckPassword(l.Password, user.Password)
	if err != nil {
		// unauthorized
		return domain.Token{}, err
	}

	accessToken, accessPayload, err := uc.jwtTokenMaker.CreateToken(
		string(user.ID),
		user.Username,
		56*time.Hour,
	)

	if err != nil {
		// internal server error
		return domain.Token{}, err
	}

	refreshToken, refreshPayload, err := uc.jwtTokenMaker.CreateToken(
		string(user.ID),
		user.Username,
		168*time.Hour,
	)
	if err != nil {
		// internal server error
		return domain.Token{}, err
	}

	createSessionReq := domain.Session{
		ID:           refreshPayload.ID,
		Username:     user.Username,
		RefreshToken: refreshToken,
		ExpiresAt:    refreshPayload.ExpiredAt,
	}

	err = uc.sessionRepo.Insert(
		ctx,
		createSessionReq,
	)

	if err != nil {
		return domain.Token{}, err
	}

	token := domain.Token{
		SessionId:             createSessionReq.ID,
		AccessToken:           accessToken,
		AccessTokenExpiresAt:  accessPayload.ExpiredAt,
		RefreshToken:          refreshToken,
		RefreshTokenExpiresAt: refreshPayload.ExpiredAt,
		User:                  user,
	}
	return token, nil
}

func (uc *AuthService) RenewAccessToken(ctx context.Context, r router.NewTokenReq) (domain.NewToken, error) {
	refreshPayload, err := uc.jwtTokenMaker.VerifyToken(r.RefreshToken)
	if err != nil {
		// Unauthorized , token yg dikrim tidak sama dg yg ada di database
		return domain.NewToken{}, domain.WrapErrorf(err, domain.ErrUnauthorized, domain.MessageUnauthorized)
	}

	session, err := uc.sessionRepo.Get(ctx, refreshPayload.ID)
	if err != nil {
		return domain.NewToken{}, domain.WrapErrorf(err, domain.ErrUnauthorized, domain.MessageUnauthorized)
	}

	if session.Username != refreshPayload.Username {
		return domain.NewToken{}, domain.WrapErrorf(err, domain.ErrUnauthorized, domain.MessageUnauthorized)
	}
	if session.RefreshToken != r.RefreshToken {
		return domain.NewToken{}, domain.WrapErrorf(err, domain.ErrUnauthorized, domain.MessageUnauthorized)
	}

	if time.Now().After(session.ExpiresAt) {
		return domain.NewToken{}, domain.WrapErrorf(err, domain.ErrUnauthorized, domain.MessageUnauthorized)
	}

	accessToken, accessTokenPayload, err := uc.jwtTokenMaker.CreateToken(
		refreshPayload.ID,
		refreshPayload.Username,
		7*time.Hour,
	)
	if err != nil {
		return domain.NewToken{}, domain.WrapErrorf(err, domain.ErrInternalServerError, domain.MessageInternalServerError)
	}

	res := domain.NewToken{
		AccessToken:          accessToken,
		AccessTokenExpiresAt: accessTokenPayload.ExpiredAt,
	}
	return res, nil
}

func (uc *AuthService) DeleteRefreshToken(ctx context.Context, d router.DeleteRefreshTokenReq) error {

	refreshPayload, err := uc.jwtTokenMaker.VerifyToken(d.RefreshToken)
	if err != nil {
		// Unauthorized , token yg dikrim tidak sama dg yg ada di database
		return err
	}

	session, err := uc.sessionRepo.Get(ctx, refreshPayload.ID)
	if err != nil {
		return err
	}
	if session.Username != refreshPayload.Username {
		return err
	}
	if session.RefreshToken != d.RefreshToken {
		return err
	}

	if time.Now().After(session.ExpiresAt) {
		return err
	}
	err = uc.sessionRepo.Delete(ctx, refreshPayload.ID)
	if err != nil {
		return err
	}

	return nil
}
