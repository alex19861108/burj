package model

import "gopkg.in/mgo.v2/bson"

type NodeStatus int

const (
	Running NodeStatus = iota
	Stopped
)

type Node struct {
	Id_   bson.ObjectId `bson:"_id"`
	Addr  string        `bson:"addr"`
	Label string        `bson:"label"`
}
