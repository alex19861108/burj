package services

import (
	"github.com/alex19861108/burj/burj_center/iris/proto"
	"github.com/alex19861108/burj/burj_center/iris/repositories"
	"gopkg.in/mgo.v2/bson"
)

type NodeService interface {
	GetAll() ([]*proto.Node, error)
	GetByID(id string) (proto.Node, error)
	GetByAddr(addr string) ([]*proto.Node, error)
	GetByLabels(labels string) ([]*proto.Node, error)
	GetByDeviceLabel(label string) ([]*proto.Node, error)
	DeleteByID(id string) error
	DeleteByLabelAndAddr(label string, addr string) error
	InsertOrUpdate(n proto.Node) (proto.Node, error)
	UpdateAddrAndLabelByID(id string, addr string, label string) (proto.Node, error)
	Clean() error
	FilterDeviceByDeviceLabel(nodes proto.Node, s string) proto.Node
}

// NewNodeService returns the default node service.
func NewNodeService(repo repositories.NodeRepository) NodeService {
	return &nodeService{
		repo: repo,
	}
}

type nodeService struct {
	repo repositories.NodeRepository
}

// GetAll returns all nodes.
func (s *nodeService) GetAll() ([]*proto.Node, error) {
	return s.repo.FindAll(nil)
}

// GetByID returns a node based on its id.
func (s *nodeService) GetByID(id string) (proto.Node, error) {
	return s.repo.FindByID(id)
}

func (s *nodeService) GetByAddr(addr string) ([]*proto.Node, error) {
	return s.repo.FindAll(bson.M{"addr.http": addr})
}

func (s *nodeService) GetByLabels(labels string) ([]*proto.Node, error) {
	return s.repo.FindAll(bson.M{"labels": labels})
}

func (s *nodeService) InsertOrUpdate(n proto.Node) (proto.Node, error) {
	return s.repo.InsertOrUpdate(n)
}

// UpdatePosterAndGenreByID updates a node's poster and genre.
func (s *nodeService) UpdateAddrAndLabelByID(id string, addr string, labels string) (proto.Node, error) {
	// update the node and return it.
	return s.repo.InsertOrUpdate(proto.Node{
		Id: string(bson.ObjectIdHex(id)),
		Addr: &proto.Addr{
			Http: addr,
		},
		Labels: labels,
	})
}

// DeleteByID deletes a node by its id.
//
// Returns true if deleted otherwise false.
func (s *nodeService) DeleteByID(id string) error {
	return s.repo.DeleteByID(bson.ObjectIdHex(id))
}

func (s *nodeService) DeleteByLabelAndAddr(label string, addr string) error {
	return s.repo.DeleteBy(bson.M{"labels": label, "addr.http": addr})
}

func (s *nodeService) Clean() error {
	return s.repo.DropCollection()
}

func (s *nodeService) GetByDeviceLabel(label string) ([]*proto.Node, error) {
	nodes, err := s.repo.FindAll(bson.M{"devices": bson.M{"$elemMatch": bson.M{"label": label}}})
	return nodes, err
}

func (s *nodeService) FilterDeviceByDeviceLabel(node proto.Node, label string) (result proto.Node) {
	var devices []*proto.Device
	for _, device := range node.Devices {
		if device.Label == "label" {
			devices = append(devices, device)
		}
	}
	node.Devices = devices
	return node
}
