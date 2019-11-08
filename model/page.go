package model

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type permission int
type WidgetsVersionMap map[string]string

const (
	Private  permission = iota // 0 私人
	Internal                   // 1 组内
	Public                     // 2 公开
)

type Page struct {
	ID             primitive.ObjectID `bson:"_id" json:"id"`                          // 页面唯一 ID
	Name           string             `bson:"name" json:"name"`                       // 页面名称
	Creater        primitive.ObjectID `bson:"creater" json:"creater"`                 // 页面创建者
	CreaterName    string             `bson:"creater_name" json:"creater_name"`       // 页面创建者
	Data           []*Components      `bson:"data" json:"data"`                       // 页面组件配置数据
	Historys       []*History         `bson:"historys" json:"historys"`               // 页面更改历史
	WidgetsVersion WidgetsVersionMap  `bson:"widgets_version" json:"widgets_version"` // 页面组件信息
	Resources      []*Resource        `bson:"resources" json:"resources"`             // 页面资源列表
	ExpiryStart    time.Time          `bson:"expiry_start" json:"expiry_start"`       // 页面上线起止时间
	ExpiryEnd      time.Time          `bson:"expiry_end" json:"expiry_end"`           // 页面上线起止时间
	Permission     permission         `bson:"permission" json:"permission"`           // 页面可见性
	Created        time.Time          `bson:"created" json:"created"`
	Updated        time.Time          `bson:"updated" json:"updated"`
}

type PgaeUpdateInfo struct {
	ID         primitive.ObjectID `bson:"_id" json:"id"`
	EditorID   primitive.ObjectID
	EditorName string
}
type PgaeUpdateData struct {
	Name        string        `json:"name"`
	Data        []*Components `json:"data"`
	Resources   []*Resource   `bson:"resources" json:"resources"`
	ExpiryStart time.Time     `json:"expiryStart"`
	ExpiryEnd   time.Time     `json:"expiryEnd"`
	Permission  permission    `json:"permission"`
}

// 页面数据
type PageData struct {
	Name           string             `json:"name" binding:"required"`
	Data           []*Components      `json:"data"`
	WidgetsVersion WidgetsVersionMap  `bson:"widgets_version" json:"widgets_version"`
	Creater        primitive.ObjectID `json:"creater"`
	CreaterName    string             `bson:"creater_name" json:"creater_name"`
	ExpiryStart    time.Time          `json:"expiry_start"`
	ExpiryEnd      time.Time          `json:"expiry_end"`
	Permission     permission         `json:"permission"`
}

// 页面信息
type PageInfo struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	Name        string             `bson:"name" json:"name"`
	Creater     primitive.ObjectID `bson:"creater" json:"creater"`
	CreaterName string             `bson:"creater_name" json:"creater_name"`
	Historys    []*History         `bson:"historys" json:"historys"`
	ExpiryStart time.Time          `bson:"expiry_start" json:"expiry_start"`
	ExpiryEnd   time.Time          `bson:"expiry_end" json:"expiry_end"`
	Permission  permission         `bson:"permission" json:"permission"`
}
type PageHistory struct {
	ID       primitive.ObjectID
	Name     string
	Historys []*History
}
type Resource struct {
	Name string                 `bson:"name" json:"name"`
	Url  string                 `bson:"url" json:"url"`
	Key  string                 `bson:"key" json:"key"`
	Info map[string]interface{} `bson:"info" json:"info"`
}
type Components struct {
	Children []*Components          `bson:"children" json:"children"`
	Name     string                 `bson:"name" json:"name"`
	ID       string                 `bson:"id" json:"id"`
	Props    map[string]interface{} `bson:"props" json:"props"`
}

type History struct {
	UpdateUserID   primitive.ObjectID `bson:"update_user_id" json:"update_user_id"`
	UpdateUsername string             `bson:"update_username" json:"update_username"`
	UpdatedTime    time.Time          `bson:"updated_time" json:"updated_time"`
}
type Widgets struct {
	Version string `bson:"version" json:"version"`
	Name    string `bson:"name" json:"name"`
	ID      string `bson:"id" json:"_id"`
}

func (p *Page) New() *Page {
	history := &History{
		UpdateUserID:   p.Creater,
		UpdateUsername: p.CreaterName,
		UpdatedTime:    time.Now(),
	}
	return &Page{
		ID:             primitive.NewObjectID(),
		Name:           p.Name,
		Creater:        p.Creater,
		CreaterName:    p.CreaterName,
		Data:           p.Data,
		WidgetsVersion: p.WidgetsVersion,
		ExpiryStart:    p.ExpiryStart,
		ExpiryEnd:      p.ExpiryEnd,
		Historys:       []*History{history},
		Created:        time.Now(),
		Updated:        time.Now(),
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

func (p *Page) GetPageDataByID(id primitive.ObjectID) (*PageData, error) {
	var pageData *PageData
	err := DB.Self.Collection("Page").FindOne(context.Background(), bson.D{{Key: "_id", Value: id}}).Decode(&pageData)
	if err != nil {
		return nil, err
	}
	return pageData, nil
}

func UpdatePage(updateInfo *PgaeUpdateInfo, updateDateMap *map[string]interface{}) error {
	_now := time.Now()
	(*updateDateMap)["updated"] = _now
	_, err := DB.Self.Collection("Page").
		UpdateOne(
			context.Background(),
			bson.D{{Key: "_id", Value: updateInfo.ID}},
			bson.M{
				"$push": bson.M{"historys": bson.M{"$each": bson.A{&History{
					UpdateUserID:   updateInfo.EditorID,
					UpdateUsername: updateInfo.EditorName,
					UpdatedTime:    _now,
				}}, "$slice": 10}},
				"$set": updateDateMap,
			},
			&options.UpdateOptions{},
		)
	if err != nil {
		return err
	}
	return nil
}
func PushPageResource(id primitive.ObjectID, resource *Resource) error {
	_, err := DB.Self.Collection("Page").
		UpdateOne(
			context.Background(),
			bson.D{{Key: "_id", Value: id}},
			bson.M{
				"$addToSet": bson.M{"resources": bson.M{"$each": bson.A{resource}, "$slice": 100}},
			},
			&options.UpdateOptions{},
		)
	if err != nil {
		return err
	}
	return nil
}
func PullPageResource(id primitive.ObjectID, key string) error {
	_, err := DB.Self.Collection("Page").
		UpdateOne(
			context.Background(),
			bson.D{{Key: "_id", Value: id}},
			bson.M{
				"$pull": bson.M{"resources": bson.M{
					"key": key,
				}},
			},
			&options.UpdateOptions{},
		)
	if err != nil {
		return err
	}
	return nil
}

type fields struct {
	Resources []*Resource `bson:"resources"`
}

func GetPageResource(id primitive.ObjectID) (error, *[]*Resource) {
	var pageData Page
	err := DB.Self.Collection("Page").FindOne(context.Background(),
		bson.D{{Key: "_id", Value: id}}, &options.FindOneOptions{
			Projection: fields{
				Resources: []*Resource{},
			},
		}).Decode(&pageData)
	if err != nil {
		return err, nil
	}

	return nil, &pageData.Resources
}

func UpdateWidgetVersion(id primitive.ObjectID, updateWidgetsMap *WidgetsVersionMap) error {
	var pageData Page
	err := DB.Self.Collection("Page").FindOne(context.Background(),
		bson.D{{Key: "_id", Value: id}}).Decode(&pageData)
	if err != nil {
		return err
	}
	for k, v := range *updateWidgetsMap {
		pageData.WidgetsVersion[k] = v
	}
	_, err = DB.Self.Collection("Page").
		UpdateOne(
			context.Background(),
			bson.D{{Key: "_id", Value: id}},
			bson.M{
				"$set": bson.M{"widgets_version": pageData.WidgetsVersion},
			},
			&options.UpdateOptions{},
		)
	if err != nil {
		return err
	}
	return nil
}
func GetWidgetVersion(id primitive.ObjectID) (error, *WidgetsVersionMap) {
	var pageData Page
	err := DB.Self.Collection("Page").FindOne(context.Background(),
		bson.D{{Key: "_id", Value: id}}).Decode(&pageData)
	if err != nil {
		return err, nil
	}

	return nil, &pageData.WidgetsVersion
}

func GetPageList(limit, skip int64) (int64, []*PageInfo, error) {
	pages := []*PageInfo{}
	filter := bson.D{{Key: "creater_name", Value: "ped"}}
	cursor, err := DB.Self.Collection("Page").
		Find(context.Background(), filter,
			&options.FindOptions{
				Limit: &limit,
				Skip:  &skip,
				Sort: bson.M{
					"created": -1,
				},
			})
	if err != nil {
		return 0, nil, err
	}
	pageTotal, err := DB.Self.Collection("Page").
		CountDocuments(context.TODO(), filter)
	if err != nil {
		return 0, nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		pageInfo := &PageInfo{}
		if err := cursor.Decode(pageInfo); err != nil {
			return 0, nil, err
		}
		pages = append(pages, pageInfo)
	}

	return pageTotal, pages, nil
}
