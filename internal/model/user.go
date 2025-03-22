package model

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID               primitive.ObjectID `bson:"_id,omitempty"`
	UID              string            `bson:"uid"`
	Username         string            `bson:"username"`
	Email            string            `bson:"email"`
	Password         string            `bson:"password"`
	DateOfBirth      *time.Time        `bson:"dateOfBirth,omitempty"`
	Gender           string            `bson:"gender,omitempty"`
	ProfilePicture   string            `bson:"profilePicture,omitempty"`
	BackgroundPicture string           `bson:"backgroundPicture,omitempty"`
	IsActive         bool              `bson:"isActive"`
	IsDeleted        bool              `bson:"isDeleted"`
	Role             string            `bson:"role"`
	Projects         []primitive.ObjectID `bson:"projects"`
	CreatedAt        time.Time         `bson:"createdAt"`
	UpdatedAt        time.Time         `bson:"updatedAt"`
} 