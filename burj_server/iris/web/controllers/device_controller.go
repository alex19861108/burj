package controllers

import (
	"github.com/alex19861108/burj/burj_center/iris/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type DeviceController struct {
	Conn *grpc.ClientConn
}

func (c *DeviceController) Get() (devices []*proto.Device, err error) {
	client := proto.NewNodeServiceClient(c.Conn)
	resp, err := client.GetDevices(context.Background(), &proto.GetDevicesRequest{})
	if err != nil || len(resp.Devices) == 0 {
		return devices, ERROR_NO_NODE_FOUND
	} else {
		return resp.Devices, nil
	}
}
