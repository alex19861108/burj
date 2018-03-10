package datamodels

import (
	"gopkg.in/mgo.v2/bson"
)

type Addr struct {
	Http string `bson:"http" json:"http"`
	Rpc  string `bson:"rpc" json:"rpc"`
}

type Node struct {
	ID      bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
	Addr    Addr          `bson:"addr" json:"addr"`
	Labels  string        `bson:"labels" json:"labels"`
	Devices []Device      `bson:"devices" json:"devices"`
}
