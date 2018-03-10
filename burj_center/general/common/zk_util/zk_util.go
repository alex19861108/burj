package zk_util

import (
	"time"

	"github.com/samuel/go-zookeeper/zk"
)

func GetConnect(zk_host string, timeout int) (conn *zk.Conn, err error) {
	conn, _, err = zk.Connect([]string{zk_host}, time.Duration(timeout)*time.Second)
	return
}

func CreateNode(conn *zk.Conn, path string) (err error) {
	_, err = conn.Create(path, nil, zk.FlagEphemeral, zk.WorldACL(zk.PermAll))
	return
}

func GetNodeList(conn *zk.Conn) (list []string, err error) {
	list, _, err = conn.Children("/")
	return
}
