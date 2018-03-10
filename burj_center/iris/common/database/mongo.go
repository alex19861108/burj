package database

import (
	"net/url"
	"strings"
	"time"

	"github.com/alex19861108/burj/burj_center/iris/config"
	"github.com/kataras/iris/core/errors"
	"gopkg.in/mgo.v2"
)

type Session struct {
	session *mgo.Session
	dbName  string
}

const (
	timeout = 10
)

func (s *Session) C(name string) *mgo.Collection {
	return s.session.DB(s.dbName).C(name)
}

func (s *Session) Close() {
	s.session.Close()
}

var (
	globalSession *mgo.Session
	dbName        string
)

func NewSession() (*Session, error) {
	if globalSession != nil {
		return &Session{session: globalSession.Clone(), dbName: dbName}, nil
	} else {
		return nil, errors.New("global session not exists!")
	}
}

func InitMgoDB(cfg config.Config) (err error) {
	u, err := url.Parse(cfg.DBMongoConfig.URL)
	if err != nil {
		panic(err)
	}

	globalSession, err = mgo.DialWithTimeout(cfg.DBMongoConfig.URL, time.Duration(timeout)*time.Second)
	if err != nil {
		panic(err)
	}
	globalSession.SetMode(mgo.Monotonic, true)
	globalSession.SetPoolLimit(4096)
	dbName = strings.Trim(u.Path, "/")

	return err
}
