package controllers

import (
	"encoding/json"

	"gopkg.in/mgo.v2/bson"

	"github.com/kataras/iris"

	"github.com/alex19861108/burj/burj_center/iris/proto"
	"github.com/alex19861108/burj/burj_center/iris/services"
)

// DeviceController is our /device controller.
// DeviceController is responsible to handle the following requests:
// All  			/device/register
// GET 				/device/login
// GET 				/device/id
type DeviceController struct {
	Ctx     iris.Context
	Service services.DeviceService
}

func (c *DeviceController) Get() (results []proto.Device) {
	return c.Service.GetAll()
}

func (c *DeviceController) GetBy(id string) (device proto.Device, err error) {
	return c.Service.GetByID(id)
}

func (c *DeviceController) AnyRegister() (device proto.Device, err error) {
	n := c.Ctx.FormValue("device")
	err = json.Unmarshal([]byte(n), &device)
	if err != nil {
		return proto.Device{}, err
	}
	return c.Service.InsertOrUpdate(device)
}

func (c *DeviceController) GetFilter() (device []proto.Device) {
	addr := c.Ctx.FormValue("addr")
	label := c.Ctx.FormValue("label")
	return c.Service.GetBy(bson.M{"addr": bson.M{"http": addr}, "label": label})
}
