package models

import (
	"beeapi/response"
	"encoding/json"
	"gorm.io/gorm"
	"time"
)

type Item struct {
	ID uint `gorm:"primaryKey"`

	Name        string // Name of the item
	Description string // Description of the item

	Price      string // Price of the item
	Quantity   string // Quantity of the item
	Retailer   string // Retailer where the item was purchased
	PartNumber string // Part number of the item

	Location string // Location where the item is stored

	CustomFields map[string]interface{} `gorm:"json"` // Custom fields for the item

	Checkouts []Checkout // List of all checkouts for this item

	CreatedAt time.Time `gorm:"<-:create"`
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (item *Item) Verify() *response.Error {
	if item.Name == "" {
		return response.CreateError(response.ErrorCodeInvalidRequest, "Name cannot be empty.", nil)
	}

	if item.Description == "" {
		return response.CreateError(response.ErrorCodeInvalidRequest, "Description cannot be empty.", nil)
	}

	return nil
}

func (item *Item) Marshal() (interface{}, error) {
	var store struct {
		ID uint `json:"id"`

		Name        string `json:"name"`
		Description string `json:"description"`

		Price      string `json:"price,omitempty"`
		Quantity   string `json:"quantity,omitempty"`
		Retailer   string `json:"retailer,omitempty"`
		PartNumber string `json:"partNumber,omitempty"`

		CustomFields map[string]interface{} `json:"customFields,omitempty"`

		Location string `json:"location,omitempty"`

		CreatedAt time.Time  `json:"createdAt"`
		UpdatedAt time.Time  `json:"updatedAt"`
		DeletedAt *time.Time `json:"deletedAt,omitempty"`
	}

	data, err := json.Marshal(item)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &store)
	if err != nil {
		return nil, err
	}

	return store, nil
}

func (item *Item) Unmarshal(data []byte) error {
	var store struct {
		Name        string `json:"name,omitempty"`
		Description string `json:"description,omitempty"`

		Price      string `json:"price,omitempty"`
		Quantity   string `json:"quantity,omitempty"`
		Retailer   string `json:"retailer,omitempty"`
		PartNumber string `json:"partNumber,omitempty"`

		CustomFields map[string]interface{} `json:"customFields,omitempty"`

		Location string `json:"location,omitempty"`
	}

	err := json.Unmarshal(data, &store)
	if err != nil {
		return err
	}

	if &store.Name != nil {
		item.Name = store.Name
	}

	if &store.Description != nil {
		item.Description = store.Description
	}

	if &store.Price != nil {
		item.Price = store.Price
	}

	if &store.Quantity != nil {
		item.Quantity = store.Quantity
	}

	if &store.Retailer != nil {
		item.Retailer = store.Retailer
	}

	if &store.PartNumber != nil {
		item.Retailer = store.Retailer
	}

	if &store.CustomFields != nil {
		// We have to do this otherwise it will delete completely omitted custom fields.
		if item.CustomFields == nil {
			item.CustomFields = make(map[string]interface{})
		}
		for key, value := range store.CustomFields {
			item.CustomFields[key] = value
		}
	}

	if &store.Location != nil {
		item.Location = store.Location
	}

	return nil
}
