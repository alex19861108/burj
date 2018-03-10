package server

import (
	"net"

	"github.com/alex19861108/burj/burj_center/iris/config"
	"github.com/alex19861108/burj/burj_center/iris/proto"
	"github.com/alex19861108/burj/burj_center/iris/services"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type RpcServer struct {
	NodeService services.NodeService
}

func (s *RpcServer) Search(ctx context.Context, req *proto.SearchNodeRequest) (*proto.SearchNodeResponse, error) {
	var units []*proto.NodeDeviceUnit
	nodes, _ := s.NodeService.GetAll()
	for _, node := range nodes {
		for _, device := range node.Devices {
			unit := &proto.NodeDeviceUnit{
				Node:   node,
				Device: device,
			}
			units = append(units, unit)
		}
	}
	return &proto.SearchNodeResponse{
		NodeDeviceUnits: units,
	}, nil
}

func InitRpcServer(cfg config.Config, service services.NodeService) {
	listener, err := net.Listen("tcp", ":"+cfg.ServerConfig.Port.Rpc)
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	proto.RegisterNodeServiceServer(s, &RpcServer{NodeService: service})
	reflection.Register(s)
	if err := s.Serve(listener); err != nil {
		panic(err)
	}
}
