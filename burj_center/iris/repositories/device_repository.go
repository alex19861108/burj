package repositories

import (
	"github.com/alex19861108/burj/burj_center/iris/common/database"
	"github.com/alex19861108/burj/burj_center/iris/errors"
	"github.com/alex19861108/burj/burj_center/iris/proto"
	"gopkg.in/mgo.v2/bson"
)

type DeviceRepository interface {
	FindAll(m bson.M) (results []proto.Device)
	FindByID(id bson.ObjectId) (device proto.Device, err error)
	InsertOrUpdate(device proto.Device) (updatedDevice proto.Device, err error)
	DeleteByID(id bson.ObjectId) (err error)
}

func NewDeviceRepository() DeviceRepository {
	return new(deviceMongoRepository)
}

type deviceMongoRepository struct {
}

const ColDevice = "device"

func (r *deviceMongoRepository) FindAll(m bson.M) (results []proto.Device) {
	session, err := database.NewSession()
	if err != nil {
		panic(err)
		return
	}
	defer session.Close()

	session.C(ColDevice).Find(m).All(&results)
	return
}

func (r *deviceMongoRepository) FindByID(id bson.ObjectId) (device proto.Device, err error) {
	session, err := database.NewSession()
	if err != nil {
		panic(err)
		return
	}
	defer session.Close()

	query := session.C(ColDevice).FindId(id)
	if n, _ := query.Count(); n > 0 {
		query.One(&device)
		return device, nil
	} else {
		return proto.Device{}, errors.ERROR_DB_NOT_FOUND_NODE
	}
}

func (r *deviceMongoRepository) InsertOrUpdate(device proto.Device) (proto.Device, error) {
	session, err := database.NewSession()
	if err != nil {
		panic(err)
		return device, err
	}
	defer session.Close()

	devices := r.FindAll(bson.M{"addr": device.Addr})
	if len(devices) == 0 {
		device.Id = bson.NewObjectId().Hex()
		err := session.C(ColDevice).Insert(device)
		return device, err
	}

	device.Id = devices[0].Id
	err = session.C(ColDevice).Update(bson.M{"_id": device.Id}, device)

	return device, err
}

func (r *deviceMongoRepository) DeleteByID(id bson.ObjectId) error {
	session, err := database.NewSession()
	if err != nil {
		panic(err)
		return err
	}
	defer session.Close()

	err = session.C(ColDevice).Remove(bson.M{"_id": id})
	return err
}
