package repository

import (
	"github.com/alex19861108/burj/burj_center/general/model"
	"gopkg.in/mgo.v2/bson"
)

const (
	C_NODE = "node"
)

func InsertNode(n *model.Node) error {
	session, err := NewSession()
	if err != nil {
		return err
	}
	defer session.Close()

	n.Id_ = bson.NewObjectId()
	return session.DB(DBName).C(C_NODE).Insert(n)
}

func Find() {

}
