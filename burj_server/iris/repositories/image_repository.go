package repositories

import (
	"github.com/alex19861108/burj/burj_center/iris/common/database"
	"github.com/alex19861108/burj/burj_center/iris/proto"
)

type ImageRepository interface {
	Select(id int64) (img proto.Image, err error)
	SelectManyBy(query string) (imgs []*proto.Image, err error)
}

func NewImageRepository() ImageRepository {
	return new(imageMysqlRepository)
}

type imageMysqlRepository struct {
}

func (r *imageMysqlRepository) Select(id int64) (img proto.Image, err error) {
	_, err = database.NewEngine().ID(id).Get(img)
	return
}

func (r *imageMysqlRepository) SelectManyBy(query string) (imgs []*proto.Image, err error) {
	err = database.NewEngine().Where(query).Find(&imgs)
	return
}
