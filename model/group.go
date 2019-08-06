package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// 分组只关乎页面的可编辑性

type Group struct {
	ID        primitive.ObjectID    `bson:"_id" json:"id"`
	Name      string                `bson:"name" json:"name"`
	Avatar    string                `bson:"avatar" json:"avatar"`
	Creater   primitive.ObjectID    `bson:"creater" json:"creater"`
	MembersID []*primitive.ObjectID `bson:"members_id" json:"members_id"`
	Created   time.Time             `bson:"created" json:"created"`
	Updated   time.Time             `bson:"updated" json:"updated"`
}
