package basemodel

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	IsDeleted bool     `json:"-"`
	Modifier  int      `json:"modifier"`
	Modified  DateTime `json:"modified"`
	Creator   int      `json:"creator"`
	Created   DateTime `json:"created"`
}

type IDModel struct {
	ID        int `json:"id" gorm:"primaryKey"`
	BaseModel `gorm:"embedded"`
}

func (s *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	s.Created = DateTime(time.Now())
	s.Modified = DateTime(time.Now())
	return nil
}

func (s *BaseModel) BeforeUpdate(tx *gorm.DB) (err error) {
	s.Modified = DateTime(time.Now())
	return nil
}

func (s *BaseModel) BeforeSave(tx *gorm.DB) (err error) {
	s.Modified = DateTime(time.Now())
	return nil
}
