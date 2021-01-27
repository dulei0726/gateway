package dao

import "github.com/jinzhu/gorm"

type Group struct {
	gorm.Model
	GroupName string `json:"group_name"`
	Status    int
	CreatorID int
}

func (g *Group) TableName() string {
	return "group"
}
