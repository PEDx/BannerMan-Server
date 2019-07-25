package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// The Page holds
type Page struct {
	ID         primitive.ObjectID `bson:"_id" json:"id"`                  // 页面唯一 ID
	Name       string             `bson:"name" json:"name"`               // 页面名称
	Creater    string             `bson:"creater" json:"creater"`         // 页面创建者
	Owner      []*string          `bson:"owner" json:"owner"`             // 页面拥有者
	Editer     string             `bson:"editer" json:"editer"`           // 页面正在编辑的人
	Data       []*Components      `bson:"data" json:"data"`               // 页面组件配置数据
	Historys   []*History         `bson:"historys" json:"historys"`       // 页面更改历史
	ExpiryDate ExpiryDate         `bson:"expiry_date" json:"expiry_date"` // 页面上线起止时间
	Created    time.Time          `bson:"created" json:"created"`
	Updated    time.Time          `bson:"updated" json:"updated"`
}
type Components struct {
	Children      []*Components          `bson:"children" json:"children"`
	Name          string                 `bson:"name" json:"name"`
	ID            string                 `bson:"id" json:"id"`
	MultContainer bool                   `bson:"mult_container" json:"multContainer"`
	Props         map[string]interface{} `bson:"props" json:"props"`
}
type History struct {
	UpdateUserID string `bson:"update_user_id" json:"update_user_id"`
	UpdatedTime  string `bson:"updated_time" json:"updated_time"`
}
type ExpiryDate struct {
	Start time.Time `bson:"start" json:"start"`
	End   time.Time `bson:"end" json:"end"`
}
