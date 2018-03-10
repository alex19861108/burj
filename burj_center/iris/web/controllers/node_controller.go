package controllers

import (
	"encoding/json"

	"github.com/alex19861108/burj/burj_center/iris/common/zk_client"
	"github.com/alex19861108/burj/burj_center/iris/config"
	"github.com/alex19861108/burj/burj_center/iris/proto"
	"github.com/alex19861108/burj/burj_center/iris/services"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

// NodeController is our /node controller.
// NodeController is responsible to handle the following requests:
// All  			/node/register
// GET 				/node/login
// GET 				/node/id
type NodeController struct {
	Ctx     iris.Context
	Service services.NodeService
	Config  config.Config
}

func (c *NodeController) Get() (results []*proto.Node, err error) {
	return c.Service.GetAll()
}

func (c *NodeController) GetBy(id string) (node proto.Node, err error) {
	return c.Service.GetByID(id)
}

func (c *NodeController) GetFilterBy(addr string) (results []*proto.Node, err error) {
	return c.Service.GetByAddr(addr)
}

func (c *NodeController) GetLabelsBy(labels string) (results []*proto.Node, err error) {
	return c.Service.GetByLabels(labels)
}

func (c *NodeController) AnyRegister() (node proto.Node, err error) {
	n := c.Ctx.FormValue("node")
	err = json.Unmarshal([]byte(n), &node)
	if err != nil {
		return proto.Node{}, err
	}
	return c.Service.InsertOrUpdate(node)
}

func (c *NodeController) AnyDeleteBy(id string) (err error) {
	return c.Service.DeleteByID(id)
}

func (c *NodeController) AnyAvailable() (node proto.Node) {
	labels := c.Ctx.FormValue("labels")
	x, _ := zk_client.SearchOneNode(c.Config, labels)
	return x
}

func (c *NodeController) BeforeActivation(b mvc.BeforeActivation) {
	// b.Dependencies().Add/Remove
	// b.Router().Use/UseGlobal/Done // and any standard API call you already know

	// 1-> Method
	// 2-> Path
	// 3-> The controller's function name to be parsed as handler
	// 4-> Any handlers that should run before the MyCustomHandler
	//b.Handle("GET", "/something/{id:long}", "MyCustomHandler", nil)
}
