package models

import "time"

type Car struct {
	Plate      string `json:"plate" gorm:"primaryKey"`
	ProducedAt time.Time
	Model  string `json:"model"`
	Color   string `json:"color"`
	OwnerRefer uint `json:"owner_refer"`
	Owner User `json:"owner" gorm:"foreignKey:OwnerRefer"`
}