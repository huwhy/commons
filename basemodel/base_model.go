package basemodel

import (
	"github.com/huwhy/commons/util/datetimes"
	"gorm.io/gorm"
)

type BaseModel struct {
	IsDeleted bool               `json:"-"`
	Modifier  int                `json:"modifier"`
	Modified  datetimes.DateTime `json:"modified"`
	Creator   int                `json:"creator"`
	Created   datetimes.DateTime `json:"created"`
}

type IDModel struct {
	ID        int `json:"id" gorm:"primaryKey"`
	BaseModel `gorm:"embedded"`
}

func (s *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	s.Created = datetimes.Now()
	s.Modified = datetimes.Now()
	return nil
}

func (s *BaseModel) BeforeUpdate(tx *gorm.DB) (err error) {
	s.Modified = datetimes.Now()
	return nil
}

func (s *BaseModel) BeforeSave(tx *gorm.DB) (err error) {
	s.Modified = datetimes.Now()
	return nil
}
