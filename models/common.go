package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

//BaseModel gorm Model
type BaseModel struct {
	ID        int64 `json:"id" gorm:"primary_key"`
	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`
	IsDelete  int64 `json:"is_delete"`
}

//BeforeCreate BeforeCreate
func (m *BaseModel) BeforeCreate(scope *gorm.Scope) (err error) {
	m.CreatedAt = time.Now().UnixNano() / 1000
	return
}

//BeforeUpdate BeforeUpdate
func (m *BaseModel) BeforeUpdate(scope *gorm.Scope) (err error) {
	m.UpdatedAt = time.Now().UnixNano() / 1000
	return
}
