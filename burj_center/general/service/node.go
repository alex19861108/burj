package service

import (
	"github.com/alex19861108/burj/burj_center/general/model"
	"github.com/alex19861108/burj/burj_center/general/repository"
)

func RegisterNode(n *model.Node) (err error) {
	return repository.InsertNode(n)
}
