package models

type Address struct {
	DoorNumber string `json:"doorNumber" bson:"doorNumber" validate:"required"`
	Street     string `json:"street,omitempty" bson:"street,omitempty"`
	City       string `json:"city,omitempty" bson:"city,omitempty"`
	PostalCode int    `json:"postalCode" bson:"postalCode" validate:"required"`
}
