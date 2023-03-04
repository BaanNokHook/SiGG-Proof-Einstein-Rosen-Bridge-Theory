package models

/*
 *
 *  SiGG-Proof-Einstein-Rosen-Bridge-Theory
 *
 */

import "time"

type BaseModel struct {
	ID        uint `gorm:"primary_key" json:"-"`
	CreatedAt uint `json:"created_at"`
	UpdatedAt uint `json:"updated_at"`
}

func (model *BaseModel) BeforeCreate() error {
	model.CreatedAt = uint(time.Now().Unix())
	return nil
}

func (model *BaseModel) BeforeUpdate() error {
	model.UpdatedAt = uint(time.Now().Unix())
	return nil
}
