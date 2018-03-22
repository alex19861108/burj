package services

import (
	"strings"

	"github.com/alex19861108/burj/burj_center/iris/proto"
	"github.com/alex19861108/burj/burj_server/iris/repositories"
)

type TagService interface {
	GetAll() ([]proto.Tag, error)
	GetByID(id int) (proto.Tag, error)
	GetByNames(tags []string) ([]*proto.Tag, error)
	GetImageTags() ([]*proto.Tag, error)
}

func NewTagService(repo repositories.TagRepository) TagService {
	return &tagService{repo: repo}
}

type tagService struct {
	repo repositories.TagRepository
}

func (s *tagService) GetAll() (results []proto.Tag, err error) {
	return
}

func (s *tagService) GetByID(id int) (mark proto.Tag, err error) {
	return
}

func (s *tagService) GetByNames(names []string) (results []*proto.Tag, err error) {
	for _, name := range names {
		name = strings.TrimSpace(name)
		tag, err := s.repo.SelectOneByMark(name)
		if err != nil {
			continue
		}
		results = append(results, &tag)
	}
	return
}

func (s *tagService) GetImageTags() ([]*proto.Tag, error) {
	return s.repo.SelectAll()
}
