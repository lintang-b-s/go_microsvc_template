// go:build wireinject
//go:build wireinject
// +build wireinject

package di

import (
	"lintang/go_hertz_template/biz/dal/db"
	"lintang/go_hertz_template/biz/service"
	"lintang/go_hertz_template/config"

	"github.com/google/wire"
	"lintang/go_hertz_template/biz/util/jwt"
)

var ProviderSet wire.ProviderSet = wire.NewSet(
	service.NewUserService,
	db.NewUserRepo,
	wire.Bind(new(service.UserRepository), new(*db.UserRepository)),
)

func InitUserService(pg *db.Postgres, cfg *config.Config) *service.UserService {
	wire.Build(
		ProviderSet,
	)
	return nil
}

var ProviderSetAuth wire.ProviderSet = wire.NewSet(
	service.NewAuthService,
	jwt.NewJWTMaker,
	db.NewUserRepo,
	db.NewSessionRepo,

	wire.Bind(new(service.UserRepository), new(*db.UserRepository)),
	wire.Bind(new(service.SessionRepo), new(*db.SessionRepository)),
	wire.Bind(new(jwt.JwtTokenMaker), new(*jwt.JWTMaker)),
)

func InitAuthService(pg *db.Postgres, cfg *config.Config) *service.AuthService {
	wire.Build(
		ProviderSetAuth,
	)
	return nil
}
