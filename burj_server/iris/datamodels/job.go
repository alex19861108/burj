package datamodels

import (
	"time"

	"github.com/alex19861108/burj/burj_center/iris/datamodels"
	"gopkg.in/mgo.v2/bson"
)

type JobStatus int

const (
	PREPARE JobStatus = iota
	RUNNING
	SUCCESS
	ERROR
)

type Apk struct {
	Path      string `bson:"path"`
	PkgName   string `bson:"pkg_name"`
	ClassName string `bson:"class_name"`
}

type SubJob struct {
	Images     []Image           `bson:"images"`
	Node       datamodels.Node   `bson:"node"`
	Device     datamodels.Device `bson:"device"`
	Status     JobStatus         `bson:"status"`
	UpdateTime time.Time         `bason:"update_time, omitempty"`
}

type Job struct {
	ID      bson.ObjectId `bson:"_id"`
	Name    string        `bson:"name"`
	SubJobs []SubJob      `bson:"sub_jobs"`
	Apk     Apk           `bson:"apk"`
	Owner   string        `bson:"owner"`
}
