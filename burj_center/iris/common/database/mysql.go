package database

import (
	"net/url"

	"github.com/alex19861108/burj/burj_center/iris/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
)

var (
	db *xorm.Engine
)

const (
	maxDBConnection = 200
)

func NewEngine() *xorm.Engine {
	return db
}

func InitMysqlDB(cfg config.Config) (err error) {
	u, err := url.Parse(cfg.DBMysqlConfig.URL)
	if err != nil {
		panic(err)
	}

	dbSrc := u.String()[len(u.Scheme)+3:]
	db, err = xorm.NewEngine(u.Scheme, dbSrc)
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(maxDBConnection)
	db.SetMapper(core.SameMapper{})

	return
}
