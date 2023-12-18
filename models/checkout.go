package models

import (
	"gorm.io/gorm"
	"time"
)

type Checkout struct {
	gorm.Model

	Returned bool `json:"returned"` // Whether the item has been returned or not

	ReturnDate *time.Time `json:"return_date"` // Date the item was returned

	ItemID uint `json:"item_id"`
	UserID uint `json:"user_id"`
}
