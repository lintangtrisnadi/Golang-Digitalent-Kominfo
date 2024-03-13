package models

import "time"

type Item struct {
	ID          uint       `gorm:"primary_key" json:"id"`
	Code        string     `json:"itemCode"`
	Description string     `json:"description"`
	Quantity    int        `json:"quantity"`
	OrderID     uint       `json:"-"`
	CreatedAt   time.Time  `json:"-"`
	UpdatedAt   time.Time  `json:"-"`
	DeletedAt   *time.Time `json:"-"`
}
