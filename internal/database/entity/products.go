package entity

import (
	"uuid"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Id        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	ShopId    uuid.UUID `gorm:"type:uuid"`
	Name      string    `gorm:"type:varchar(100);not null"`
	Stock     uint32    `gorm:"not null;default:0"`
	Price     uint32    `gorm:"not null;default:0"`
	Available string    `gorm:"type:varchar(15);not null"`
	ImgPath   string    `gorm:"type:text"`

	Shop Shop `gorm:"foreignKey:ShopId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"company"`
}
