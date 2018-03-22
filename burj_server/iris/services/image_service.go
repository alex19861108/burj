package services

import (
	"strconv"

	"github.com/alex19861108/burj/burj_center/iris/proto"
	"github.com/alex19861108/burj/burj_server/iris/repositories"
)

type ImageService interface {
	GetAll() ([]proto.Image, error)
	GetByID(id int) (proto.Image, error)
	GetByTags(tags []*proto.Tag) ([]*proto.Image, error)
}

func NewImageService(repo repositories.ImageRepository) ImageService {
	return &imageService{repo: repo}
}

type imageService struct {
	repo repositories.ImageRepository
}

func (s *imageService) GetAll() (results []proto.Image, err error) {
	return
}

func (s *imageService) GetByID(id int) (image proto.Image, err error) {
	return
}

func (s *imageService) GetByTags(tags []*proto.Tag) (results []*proto.Image, err error) {
	sql := ""
	for id, tag := range tags {
		if id != 0 {
			sql = sql + " or "
		}
		sql = sql + `Mark LIKE '%"` + strconv.FormatInt(tag.Id, 10) + `"%'`
	}
	return s.repo.SelectManyBy(sql)
}
