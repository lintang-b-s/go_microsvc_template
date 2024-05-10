package dal

import (
	"lintang/go_hertz_template/biz/dal/db"
	"lintang/go_hertz_template/config"
)

func InitPg(cfg *config.Config) *db.Postgres {
	pg := db.NewPostgres(cfg)
	return pg
}
