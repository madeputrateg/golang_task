package entity

import (
	"uuid"

	"gorm.io/gorm"
)

type Shop struct {
	gorm.Model
	Id       uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Name     string    `gorm:"type:varchar(100);not null"`
	Location string    `gorm:"type:varchar(100);not null"`
	ImgPath  string    `gorm:"type:text"`
}
