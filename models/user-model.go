package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FullName    string             `json:"fullName" bson:"fullName" validate:"required"`
	Email       string             `json:"email" bson:"email" validate:"required"`
	Address     Address            `json:"address" bson:"address"`
	CreatedDate primitive.DateTime `json:"createdDate,omitempty" bson:"createdDate,omitempty"`
}
