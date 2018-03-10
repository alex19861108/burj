package datamodels

import "time"

type Tag struct {
	Id   int64
	Mark string `xorm:"varchar(100) notnull" json:"mark"`
}

func (i *Tag) TableName() string {
	return "Marks"
}

type Image struct {
	Id           string    `xorm:"varchar(50) notnull pk" json:"id"`
	Mark         string    `xorm:"varchar(50)" json:"mark"`
	Path         string    `xorm:"varchar(100) notnull" json:"path"`
	LastUpdateBy string    `xorm:"varchar(40) notnull" json:"last_update_by"`
	LastUpdate   time.Time `xorm:"date" json:"last_update"`
}

func (i *Image) TableName() string {
	return "Pictures"
}

type User struct {
	Id       int64
	UserName string `xorm:"varchar(20)" json:"user_name"`
	PassWord string `xorm:"varchar(50)" json:"password"`
	Role     string `xorm:"varchar(20)" json:"role"`
}
