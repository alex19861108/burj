package repositories

import (
	"github.com/alex19861108/burj/burj_center/iris/common/database"
	"github.com/alex19861108/burj/burj_center/iris/proto"
)

type TagRepository interface {
	Select(id int64) (img proto.Tag, err error)
	SelectOneByMark(query string) (img proto.Tag, err error)
}

func NewTagRepository() TagRepository {
	return new(tagMysqlRepository)
}

type tagMysqlRepository struct{}

func (r *tagMysqlRepository) Select(id int64) (tag proto.Tag, err error) {
	_, err = database.NewEngine().ID(id).Get(tag)
	return
}

func (r *tagMysqlRepository) SelectOneByMark(query string) (tag proto.Tag, err error) {
	_, err = database.NewEngine().Where("Mark=?", query).Get(&tag)
	if err != nil {
		panic(err)
	}
	return
}
