package services

import (
	"github.com/alex19861108/burj/burj_center/iris/proto"
	"github.com/alex19861108/burj/burj_center/iris/repositories"
	"gopkg.in/mgo.v2/bson"
)

type DeviceService interface {
	GetAll() []proto.Device
	GetByID(id string) (proto.Device, error)
	GetBy(m bson.M) []proto.Device
	DeleteByID(id string) error
	InsertOrUpdate(n proto.Device) (proto.Device, error)
}

func NewDeviceService(repo repositories.DeviceRepository) DeviceService {
	return &deviceService{
		repo: repo,
	}
}

type deviceService struct {
	repo repositories.DeviceRepository
}

func (s *deviceService) GetAll() []proto.Device {
	return s.repo.FindAll(nil)
}

func (s *deviceService) GetByID(id string) (proto.Device, error) {
	return s.repo.FindByID(bson.ObjectIdHex(id))
}

func (s *deviceService) InsertOrUpdate(p proto.Device) (proto.Device, error) {
	return s.repo.InsertOrUpdate(p)
}

func (s *deviceService) DeleteByID(id string) error {
	return s.repo.DeleteByID(bson.ObjectIdHex(id))
}

func (s *deviceService) GetBy(m bson.M) []proto.Device {
	return s.repo.FindAll(m)
}
