package services

import (
	"log"
	"math"
	"strconv"
	"time"

	"github.com/alex19861108/burj/burj_center/iris/proto"
	"github.com/alex19861108/burj/burj_server/iris/repositories"
	"gopkg.in/mgo.v2/bson"
)

type JobService interface {
	GetAll() []proto.Job
	GetByID(id string) (proto.Job, error)
	GetByOwner(addr string) []proto.Job
	DeleteByID(id string) error
	InsertOrUpdate(j proto.Job) (proto.Job, error)
	PacketImagesToDevices(images []*proto.Image, units []*proto.NodeDeviceUnit) (subJobs []*proto.SubJob)
}

func NewJobService(repo repositories.JobRepository) JobService {
	return &jobService{
		repo: repo,
	}
}

type jobService struct {
	repo repositories.JobRepository
}

func (s *jobService) GetAll() []proto.Job {
	return s.repo.FindAll(nil)
}

func (s *jobService) GetByID(id string) (proto.Job, error) {
	return s.repo.FindByID(id)
}

func (s *jobService) GetByOwner(owner string) []proto.Job {
	return s.repo.FindAll(bson.M{"owner": owner})
}

func (s *jobService) InsertOrUpdate(j proto.Job) (proto.Job, error) {
	return s.repo.InsertOrUpdate(j)
}

func (s *jobService) DeleteByID(id string) error {
	return s.repo.DeleteByID(id)
}

func (s *jobService) PacketImagesToDevices(images []*proto.Image, units []*proto.NodeDeviceUnit) (subJobs []*proto.SubJob) {
	deviceCount := len(units)

	part := len(images) / deviceCount
	for id, unit := range units {
		begin := id * part
		end := int(math.Min(float64((id+1)*part), float64(len(images))))
		//protoTime, _ := ptypes.TimestampProto(time.Now())
		subJob := &proto.SubJob{
			Id:         strconv.Itoa(id),
			Images:     images[begin:end],
			Node:       unit.Node,
			Device:     unit.Device,
			Status:     proto.Status_PREPARE,
			UpdateTime: time.Now().String(),
		}
		log.Printf("%#v\n", subJob)
		subJobs = append(subJobs, subJob)
	}
	return subJobs
}
