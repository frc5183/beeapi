package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model

	Name        string `json:"name"`        // Name of the item
	Description string `json:"description"` // Description of the item

	Price      int    `json:"price"`       // Price of the item
	Quantity   int    `json:"quantity"`    // Quantity of the item
	Retailer   string `json:"retailer"`    // Retailer where the item was purchased
	PartNumber string `json:"part_number"` // Part number of the item

	Location string `json:"location"` // Location where the item is stored

	Checkouts []Checkout `json:"checkouts"` // List of all checkouts for this item
}
