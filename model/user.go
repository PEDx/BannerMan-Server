package model

import (
	"context"
	"time"

	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	validator "gopkg.in/go-playground/validator.v9"
)

type role int

const (
	Admin   role = iota // 超级用户 (所有权限)
	Manager             // 管理员 (管理用户, 管理分组, 管理控件, 管理主题)
	Editor              // 普通编辑器用户 (编辑发布页面)
	Guest               // 游客 (只可查看页面)
)

// 页面编辑权  由 Permission 和 Group 控制
// 编辑器管理权 由 Role 控制

// The User holds
type User struct {
	ID         primitive.ObjectID    `bson:"_id" json:"id"`
	Avatar     string                `bson:"avatar" json:"avatar"`
	Username   string                `bson:"username" json:"username" validate:"min=1,max=32"`
	Password   string                `bson:"password" json:"password" validate:"min=5,max=128"`
	Email      string                `bson:"email" json:"email"`
	Role       role                  `bson:"role" json:"role"`
	Groups     []*primitive.ObjectID `bson:"groups" json:"groups"`
	Phone      string                `bson:"phone" json:"phone"`
	OwnWidgets []*Widget             `bson:"own_widgets" json:"own_widgets"`
	Created    time.Time             `bson:"created" json:"created"`
	Updated    time.Time             `bson:"updated" json:"updated"`
}

type UserInfo struct {
	ID         primitive.ObjectID    `bson:"_id" json:"id"`
	Avatar     string                `bson:"avatar" json:"avatar"`
	Username   string                `bson:"username" json:"username" `
	Phone      string                `bson:"phone" json:"phone"`
	Groups     []*primitive.ObjectID `bson:"groups" json:"groups"`
	Role       string                `bson:"role" json:"role"`
	Email      string                `bson:"email" json:"email"`
	OwnWidgets []*Widget             `bson:"own_widgets" json:"own_widgets"`
}
type Widget struct {
	WidgetName string `bson:"widget_name" json:"widget_name"`
	Name       string `bson:"name" json:"name"`
	GroupName  string `bson:"group_name" json:"group_name"`
}

func (u *User) New() *User {
	return &User{
		ID:          primitive.NewObjectID(),
		Name:        u.Name,
		Username:    u.Username,
		Email:       u.Email,
		Avatar:      u.Avatar,
		Password:    u.Password,
		Phone:       u.Phone,
		Role:        u.Role,
		IsGroupUser: u.IsGroupUser,
		MembersID:   u.MembersID,
		OwnWidgets:  u.OwnWidgets,
		Created:     time.Now(),
		Updated:     time.Now(),
	}
}

func (u *User) CreateUser() error {

	if _, err := DB.Self.Collection("User").InsertOne(context.Background(), u); err != nil {
		return err
	}
	return nil
}
func (u *User) DeleteUserByID(id primitive.ObjectID) error {

	if _, err := DB.Self.Collection("User").DeleteOne(context.Background(), bson.D{{Key: "_id", Value: id}}); err != nil {
		return err
	}
	return nil
}

func (u *User) GetUserByIDs(ids *[]primitive.ObjectID) error {

	if _, err := DB.Self.Collection("User").Find(context.Background(), bson.D{{
		Key: "_id",
		Value: bson.D{{
			Key:   "$in",
			Value: ids,
		}},
	}}); err != nil {
		return err
	}
	return nil
}
func (u *User) GetUserByUsername(username string) error {

	if _, err := DB.Self.Collection("User").Find(context.Background(), bson.D{{
		Key: "username",
		Value: bson.D{{
			Key:   "$in",
			Value: username,
		}},
	}}); err != nil {
		return err
	}
	return nil
}
func GetUserList(limit, skip int64) ([]*UserInfo, error) {
	users := []*UserInfo{}
	cursor, err := DB.Self.Collection("User").
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
		userInfo := &UserInfo{}
		if err := cursor.Decode(userInfo); err != nil {
			return nil, err
		}
		users = append(users, userInfo)
	}

	return users, nil
}

func CountUser() string {
	total, err := DB.Self.Collection("User").CountDocuments(context.Background(), bson.D{{}}, &options.CountOptions{})
	if err != nil {
		return "0"
	}
	return strconv.Itoa(int(total))
}

func (u *User) UpdateUser() *User {
	u.Updated = time.Now()
	result := DB.Self.Collection("User").
		FindOneAndReplace(context.Background(),
			bson.D{{Key: "_id", Value: u.ID}},
			u,
			&options.FindOneAndReplaceOptions{},
		)
	if result != nil {
		return u
	}
	return nil
}

// Validate the fields.
func (u *User) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
