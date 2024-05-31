package db

import (
	"context"
	"database/sql"
	"fmt"
	"lintang/go_hertz_template/config"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

type Mysql struct {
	Conn *sql.DB
}

func dsn(cfg *config.Config) string {
	zap.L().Info(fmt.Sprintf("database name: %s", cfg.Mysql.Database))

	return fmt.Sprintf("%s:%s@tcp(%s)/%s", cfg.Mysql.Username, cfg.Mysql.Password, cfg.Mysql.Host, cfg.Mysql.Database)

}

func NewMYSQL(cfg *config.Config) *Mysql {

	db, err := sql.Open("mysql", dsn(cfg))
	if err != nil {
		zap.L().Fatal("sql.Open (NewMysql)", zap.Error(err))
	}

	// q := dsn.Query()
	// q.Add("sslmode", "disable")

	// dsn.RawQuery = q.Encode()

	// dbConfig, err := pgxpool.ParseConfig(dsn.String())
	// dbConfig.MaxConns = 10
	// dbConfig.MinConns = 2
	// pool, err := pgxpool.NewWithConfig(context.Background(),  dbConfig)
	// if err != nil {
	// 	zap.L().Fatal("pgxpool connect", zap.Error(err))
	// }

	if err := db.PingContext(context.Background()); err != nil {
		hlog.Fatal("db.PingContext", zap.Error(err))
	}
	return &Mysql{Conn: db}
}

func (mysql *Mysql) ClosePostgres(ctx context.Context) {
	zap.L().Info("closing postgres gracefully")
	mysql.Conn.Close()
}
