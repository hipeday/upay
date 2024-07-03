package repository

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hipeday/upay/config"
	"github.com/jmoiron/sqlx"
)

func InitMySQL(cfg config.Config) (*sqlx.DB, error) {
	mysqlCfg := cfg.Database.MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		mysqlCfg.Username, mysqlCfg.Password, mysqlCfg.Host, mysqlCfg.Port, mysqlCfg.Database)
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func InitRabbitMQ(cfg config.Config) {

}