package basemodel

import (
	"gorm.io/gorm"
	"huwhy.cn/commons/util/datetimes"
)

type BaseModel struct {
	IsDeleted bool               `json:"-"`
	Modifier  uint               `json:"modifier"`
	Modified  datetimes.DateTime `json:"modified"`
	Creator   uint               `json:"creator"`
	Created   datetimes.DateTime `json:"created"`
}

type IDModel struct {
	ID        uint `json:"id" gorm:"primaryKey"`
	BaseModel `gorm:"embedded"`
}

func (s *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	s.Created = datetimes.Now()
	s.Modified = datetimes.Now()
	return nil
}

func (s *BaseModel) BeforeSave(tx *gorm.DB) (err error) {
	s.Modified = datetimes.Now()
	return nil
}
