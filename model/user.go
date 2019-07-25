package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// The User holds
type User struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	Name        string             `bson:"name" json:"name"`
	Avatar      string             `bson:"avatar" json:"avatar"`
	UserName    string             `bson:"username" json:"username"`
	Password    string             `bson:"password" json:"password"`
	Email       string             `bson:"email" json:"email"`
	Role        string             `bson:"role" json:"role"`
	IsGroupUser bool               `bson:"is_group_user" json:"is_group_user"` // 是否是组用户
	MembersID   []*string          `bson:"members_id" json:"members_id"`
	Phone       string             `bson:"phone" json:"phone"`
	OwnWidgets  []*Widget          `bson:"own_widgets" json:"own_widgets"`
	Created     time.Time          `bson:"created" json:"created"`
	Updated     time.Time          `bson:"updated" json:"updated"`
}

type Widget struct {
	WidgetName string `bson:"widget_name" json:"widget_name"`
	Name       string `bson:"name" json:"name"`
	GroupName  string `bson:"group_name" json:"group_name"`
}
