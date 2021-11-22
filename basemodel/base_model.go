package basemodel

import (
	"github.com/huwhy/commons/util/datetimes"
	"gorm.io/gorm"
)

type BaseModel struct {
	IsDeleted bool                `json:"-"`
	Modifier  int64               `json:"modifier"`
	Modified  *datetimes.DateTime `json:"modified"`
	Creator   int64               `json:"creator"`
	Created   *datetimes.DateTime `json:"created"`
}

type IDModel struct {
	ID        int64 `json:"id" gorm:"primaryKey"`
	BaseModel `gorm:"embedded"`
}

func (s *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	if s.Created == nil {
		s.Created = datetimes.Now2()
	}
	if s.Modified == nil {
		s.Modified = datetimes.Now2()
	}
	return nil
}

func (s *BaseModel) BeforeUpdate(tx *gorm.DB) (err error) {
	if s.Modified == nil {
		s.Modified = datetimes.Now2()
	}
	return nil
}

func (s *BaseModel) BeforeSave(tx *gorm.DB) (err error) {
	if s.Modified == nil {
		s.Modified = datetimes.Now2()
	}
	return nil
}
