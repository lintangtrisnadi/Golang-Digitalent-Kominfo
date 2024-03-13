package models

import "time"

type Order struct {
	ID           uint       `gorm:"primary_key" json:"id"`
	CustomerName string     `json:"customerName"`
	OrderedAt    time.Time  `json:"orderedAt"`
	Items        []Item     `json:"items"`
	CreatedAt    time.Time  `json:"-"`
	UpdatedAt    time.Time  `json:"-"`
	DeletedAt    *time.Time `json:"-"`
}
