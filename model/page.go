package model

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type permission int

const (
	Private  permission = iota // 0 私人
	Internal                   // 1 组内
	Public                     // 2 公开
)

type Page struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`                    // 页面唯一 ID
	Name        string             `bson:"name" json:"name"`                 // 页面名称
	Creater     primitive.ObjectID `bson:"creater" json:"creater"`           // 页面创建者
	CreaterName string             `bson:"creater_name" json:"creater_name"` // 页面创建者
	Data        []*Components      `bson:"data" json:"data"`                 // 页面组件配置数据
	Historys    []*History         `bson:"historys" json:"historys"`         // 页面更改历史
	ExpiryStart time.Time          `bson:"expiry_start" json:"expiry_start"` // 页面上线起止时间
	ExpiryEnd   time.Time          `bson:"expiry_end" json:"expiry_end"`     // 页面上线起止时间
	Permission  permission         `bson:"permission" json:"permission"`     // 页面可见性
	Created     time.Time          `bson:"created" json:"created"`
	Updated     time.Time          `bson:"updated" json:"updated"`
}

type PgaeUpdateInfo struct {
	ID         primitive.ObjectID
	EditorID   primitive.ObjectID
	EditorName string
}
type PgaeUpdateData struct {
	Data        []*Components `bson:"data" json:"data"`
	ExpiryStart time.Time     `bson:"expiry_start" json:"expiryStart"`
	ExpiryEnd   time.Time     `bson:"expiry_end" json:"expiryEnd"`
	Permission  permission    `bson:"permission" json:"permission"`
}

type PageInfo struct {
	Name        string             `json:"name" binding:"required"`
	Creater     primitive.ObjectID `json:"creater"`
	CreaterName string             `json:"creater_name" binding:"required"`
	ExpiryStart time.Time          `json:"expiry_start"`
	ExpiryEnd   time.Time          `json:"expiry_end"`
	Permission  permission         `json:"permission"`
}
type PageData struct {
	ID          primitive.ObjectID
	Name        string
	Creater     primitive.ObjectID
	CreaterName string
	Data        []*Components
	Historys    []*History
	ExpiryStart time.Time
	ExpiryEnd   time.Time
	Permission  permission
}
type PageHistory struct {
	ID       primitive.ObjectID
	Name     string
	Historys []*History
}
type Components struct {
	Children      []*Components          `bson:"children" json:"children"`
	Name          string                 `bson:"name" json:"name"`
	ID            string                 `bson:"id" json:"id"`
	MultContainer bool                   `bson:"mult_container" json:"multContainer"`
	Props         map[string]interface{} `bson:"props" json:"props"`
}

type History struct {
	UpdateUserID   primitive.ObjectID `bson:"update_user_id" json:"update_user_id"`
	UpdateUsername string             `bson:"update_username" json:"update_username"`
	UpdatedTime    time.Time          `bson:"updated_time" json:"updated_time"`
}

func (p *Page) New() *Page {
	history := &History{
		UpdateUserID:   p.Creater,
		UpdateUsername: p.CreaterName,
		UpdatedTime:    time.Now(),
	}
	return &Page{
		ID:          primitive.NewObjectID(),
		Name:        p.Name,
		Creater:     p.Creater,
		CreaterName: p.CreaterName,
		Data:        p.Data,
		ExpiryStart: p.ExpiryStart,
		ExpiryEnd:   p.ExpiryEnd,
		Historys:    []*History{history},
		Created:     time.Now(),
		Updated:     time.Now(),
	}
}
func (p *Page) CreatePage() error {

	if _, err := DB.Self.Collection("Page").InsertOne(context.Background(), p); err != nil {
		return err
	}
	return nil
}

func (p *Page) DeletePageByID(id primitive.ObjectID) error {

	if _, err := DB.Self.Collection("Page").DeleteOne(context.Background(), bson.D{{Key: "_id", Value: id}}); err != nil {
		return err
	}
	return nil
}

func (p *Page) PushPageHistory(pageID, userID primitive.ObjectID, name string) *Page {
	result := DB.Self.Collection("Page").FindOneAndUpdate(
		context.Background(),
		bson.D{{Key: "_id", Value: pageID}},
		bson.M{"$push": bson.M{"Historys": bson.M{"$each": bson.A{&History{
			UpdateUserID:   userID,
			UpdateUsername: name,
			UpdatedTime:    time.Now(),
		}}}}})
	if result != nil {
		return p
	}
	return nil
}

func (p *Page) GetPageInfoByID(id primitive.ObjectID) (*PageInfo, error) {
	var pageInfo *PageInfo
	err := DB.Self.Collection("Page").FindOne(context.Background(), bson.D{{Key: "_id", Value: id}}).Decode(&pageInfo)
	if err != nil {
		return nil, err
	}
	return pageInfo, nil
}

func UpdatePage(updateInfo *PgaeUpdateInfo, updateDateMap *map[string]interface{}) error {
	// update := bson.D{}
	// for k, v := range *updateDateMap {
	// 	update = append(update, bson.E{
	// 		Key:   k,
	// 		Value: v,
	// 	})
	// }

	_, err := DB.Self.Collection("Page").
		UpdateOne(
			context.Background(),
			bson.D{{Key: "_id", Value: updateInfo.ID}},
			bson.M{
				"$push": bson.M{"historys": bson.M{"$each": bson.A{&History{
					UpdateUserID:   updateInfo.EditorID,
					UpdateUsername: updateInfo.EditorName,
					UpdatedTime:    time.Now(),
				}}}},
				"$set": updateDateMap,
			},
			&options.UpdateOptions{},
		)
	if err != nil {
		return err
	}
	return nil
}

func GetPageList(limit, skip int64) ([]*PageInfo, error) {
	pages := []*PageInfo{}
	cursor, err := DB.Self.Collection("Page").
		Find(context.Background(), bson.D{},
			&options.FindOptions{
				Limit: &limit,
				Skip:  &skip,
			})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		pageInfo := &PageInfo{}
		if err := cursor.Decode(pageInfo); err != nil {
			return nil, err
		}
		pages = append(pages, pageInfo)
	}

	return pages, nil
}
