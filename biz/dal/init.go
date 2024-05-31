package dal

import (
	"lintang/go_hertz_template/biz/dal/db"
	"lintang/go_hertz_template/config"
)

func InitMysql(cfg *config.Config) *db.Mysql {
	pg := db.NewMYSQL(cfg)
	return pg
}
