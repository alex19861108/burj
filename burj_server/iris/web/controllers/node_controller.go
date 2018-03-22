package controllers

import (
	"github.com/alex19861108/burj/burj_center/iris/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type NodeController struct {
	Conn *grpc.ClientConn
}

func (c *NodeController) Get() (nodes []*proto.Node, err error) {
	client := proto.NewNodeServiceClient(c.Conn)
	resp, err := client.GetNodes(context.Background(), &proto.GetNodesRequest{})
	if err != nil || len(resp.Nodes) == 0 {
		return nodes, ERROR_NO_NODE_FOUND
	} else {
		return resp.Nodes, nil
	}
}

func (c *NodeController) GetBy(id string) (nodes []*proto.Node, err error) {
	client := proto.NewNodeServiceClient(c.Conn)
	resp, err := client.GetNodes(context.Background(), &proto.GetNodesRequest{Id: id})
	if err != nil || len(resp.Nodes) == 0 {
		return nodes, ERROR_NO_NODE_FOUND
	} else {
		return resp.Nodes, nil
	}
}
