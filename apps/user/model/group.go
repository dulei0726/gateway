package model

import (
	"time"
)

type Group struct {
	ID        int `gorm:"primary_key"`
	Name string
	Status    int
	CreatorID int
	User      User `gorm:"foreignKey:CreatorID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (g *Group) TableName() string {
	return "group"
}
