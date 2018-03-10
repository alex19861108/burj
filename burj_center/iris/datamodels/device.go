package datamodels

import "gopkg.in/mgo.v2/bson"

type CPU struct {
	Abi       string `bson:"abi" json:"abi"`
	Abilist   string `bson:"abilist" json:"abilist"`
	Abilist32 string `bson:"abilist32" json:"abilist32"`
	Abilist64 string `bson:"abilist64" json:"abilist64"`
}

type Memory struct {
	Total     string `bson:"total" json:"total"`
	Used      string `bson:"used" json:"used"`
	Available string `bson:"available" json:"available"`
}

type Device struct {
	ID            bson.ObjectId `bson:"_id,omitempty" json:"id, omitempty"`
	Addr          string        `bson:"addr" json:"addr"`
	Label         string        `bson:"label" json:"label"`
	CPU           CPU           `bson:"cpu" json:"cpu"`
	Memory        Memory        `bson:"memory" json:"memory"`
	Model         string        `bson:"model" json:"model"`
	Brand         string        `bson:"brand" json:"brand"`
	Name          string        `bson:"name" json:"name"`
	Device        string        `bson:"device" json:"device"`
	Board         string        `bson:"board" json:"board"`
	Locale        string        `bson:"locale" json:"locale"`
	FirstApiLevel string        `bson:"first_api_level" json:"first_api_level"`
	Manufacturer  string        `bson:"manufacturer" json:"manufacturer"`
	Cuptsm        string        `bson:"cuptsm" json:"cuptsm"`
	Serial        string        `bson:"serial" json:"serial"`
}
