package repository

import (
	"net/url"
	"strings"
	"time"

	"github.com/alex19861108/burj/burj_center/general/config"
	"github.com/alex19861108/burj/burj_center/general/types"
	"gopkg.in/mgo.v2"
)

var (
	GlobalSession *mgo.Session
	DBName        string
)

func NewSession() (*mgo.Session, error) {
	if GlobalSession != nil {
		return GlobalSession.Clone(), nil
	} else {
		return nil, types.ERROR_DB_SESSION_NOT_INITIALIZED
	}
}

func InitDB(cfg config.Config) (err error) {
	GlobalSession, err = mgo.DialWithTimeout(cfg.DBMongoConfig.URL, time.Duration(10)*time.Second)
	if err != nil {
		panic(err)
	}
	GlobalSession.SetMode(mgo.Monotonic, true)
	GlobalSession.SetPoolLimit(4096)
	u, err := url.Parse(cfg.DBMongoConfig.URL)
	DBName = strings.Trim(u.Path, "/")

	return
}
