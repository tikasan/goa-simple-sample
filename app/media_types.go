// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "goa simple sample": Application Media Types
//
// Command:
// $ goagen
// --design=github.com/pei0804/goa-simple-sample/design
// --out=$(GOPATH)/src/github.com/pei0804/goa-simple-sample
// --version=v1.3.1

package app

import (
	"github.com/goadesign/goa"
)

// celler account (default view)
//
// Identifier: application/vnd.account+json; view=default
type Account struct {
	// メールアドレス
	Email string `form:"email" json:"email" xml:"email"`
	// id
	ID int `form:"id" json:"id" xml:"id"`
	// 名前
	Name string `form:"name" json:"name" xml:"name"`
}

// Validate validates the Account media type instance.
func (mt *Account) Validate() (err error) {

	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if mt.Email == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "email"))
	}
	return
}

// AccountCollection is the media type for an array of Account (default view)
//
// Identifier: application/vnd.account+json; type=collection; view=default
type AccountCollection []*Account

// Validate validates the AccountCollection media type instance.
func (mt AccountCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// example (default view)
//
// Identifier: application/vnd.accountmedia+json; view=default
type Accountmedia struct {
	Data   []*Account `form:"data" json:"data" xml:"data"`
	Status int        `form:"status" json:"status" xml:"status"`
}

// Validate validates the Accountmedia media type instance.
func (mt *Accountmedia) Validate() (err error) {
	if mt.Data == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "data"))
	}

	for _, e := range mt.Data {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// example (default view)
//
// Identifier: application/vnd.article+json; view=default
type Article struct {
	Data     []*Data `form:"data" json:"data" xml:"data"`
	Response *OK     `form:"response" json:"response" xml:"response"`
}

// Validate validates the Article media type instance.
func (mt *Article) Validate() (err error) {
	if mt.Data == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "data"))
	}
	if mt.Response == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "response"))
	}
	for _, e := range mt.Data {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// celler bottles (default view)
//
// Identifier: application/vnd.bottle+json; view=default
type Bottle struct {
	Account *Account `form:"account" json:"account" xml:"account"`
	// id
	ID int `form:"id" json:"id" xml:"id"`
	// ボトル名
	Name string `form:"name" json:"name" xml:"name"`
	// 数量
	Quantity int `form:"quantity" json:"quantity" xml:"quantity"`
}

// Validate validates the Bottle media type instance.
func (mt *Bottle) Validate() (err error) {

	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}

	if mt.Account == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "account"))
	}
	if mt.Account != nil {
		if err2 := mt.Account.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// celler bottles (relation view)
//
// Identifier: application/vnd.bottle+json; view=relation
type BottleRelation struct {
	Account    *Account    `form:"account" json:"account" xml:"account"`
	Categories []*Category `form:"categories" json:"categories" xml:"categories"`
	// id
	ID int `form:"id" json:"id" xml:"id"`
	// ボトル名
	Name string `form:"name" json:"name" xml:"name"`
	// 数量
	Quantity int `form:"quantity" json:"quantity" xml:"quantity"`
}

// Validate validates the BottleRelation media type instance.
func (mt *BottleRelation) Validate() (err error) {

	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}

	if mt.Account == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "account"))
	}
	if mt.Categories == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "categories"))
	}
	if mt.Account != nil {
		if err2 := mt.Account.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	for _, e := range mt.Categories {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// BottleCollection is the media type for an array of Bottle (default view)
//
// Identifier: application/vnd.bottle+json; type=collection; view=default
type BottleCollection []*Bottle

// Validate validates the BottleCollection media type instance.
func (mt BottleCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// BottleCollection is the media type for an array of Bottle (relation view)
//
// Identifier: application/vnd.bottle+json; type=collection; view=relation
type BottleRelationCollection []*BottleRelation

// Validate validates the BottleRelationCollection media type instance.
func (mt BottleRelationCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// example (default view)
//
// Identifier: application/vnd.bottlemedia+json; view=default
type Bottlemedia struct {
	Data   []*Bottle `form:"data" json:"data" xml:"data"`
	Status int       `form:"status" json:"status" xml:"status"`
}

// Validate validates the Bottlemedia media type instance.
func (mt *Bottlemedia) Validate() (err error) {
	if mt.Data == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "data"))
	}

	for _, e := range mt.Data {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// celler account (default view)
//
// Identifier: application/vnd.category+json; view=default
type Category struct {
	// id
	ID int `form:"id" json:"id" xml:"id"`
	// 名前
	Name string `form:"name" json:"name" xml:"name"`
}

// Validate validates the Category media type instance.
func (mt *Category) Validate() (err error) {

	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	return
}

// example (default view)
//
// Identifier: application/vnd.categorymedia+json; view=default
type Categorymedia struct {
	Data   []*Category `form:"data" json:"data" xml:"data"`
	Status int         `form:"status" json:"status" xml:"status"`
}

// Validate validates the Categorymedia media type instance.
func (mt *Categorymedia) Validate() (err error) {
	if mt.Data == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "data"))
	}

	for _, e := range mt.Data {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// Error media type (default view)
//
// Identifier: application/vnd.error+json; view=default
type Error struct {
	Response *ErrorValue `form:"response" json:"response" xml:"response"`
}

// Validate validates the Error media type instance.
func (mt *Error) Validate() (err error) {
	if mt.Response == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "response"))
	}
	if mt.Response != nil {
		if err2 := mt.Response.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// example (default view)
//
// Identifier: application/vnd.integer+json; view=default
type Integer struct {
	// id
	ID int `form:"id" json:"id" xml:"id"`
}

// example (default view)
//
// Identifier: application/vnd.message+json; view=default
type Message struct {
	// メッセージ
	Message string `form:"message" json:"message" xml:"message"`
}

// Validate validates the Message media type instance.
func (mt *Message) Validate() (err error) {
	if mt.Message == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "message"))
	}
	return
}

// example (default view)
//
// Identifier: application/vnd.user+json; view=default
type User struct {
	// メールアドレス
	Email string `form:"email" json:"email" xml:"email"`
	// id
	ID int `form:"id" json:"id" xml:"id"`
	// 名前
	Name string `form:"name" json:"name" xml:"name"`
}

// Validate validates the User media type instance.
func (mt *User) Validate() (err error) {

	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if mt.Email == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "email"))
	}
	return
}

// example (tiny view)
//
// Identifier: application/vnd.user+json; view=tiny
type UserTiny struct {
	// id
	ID int `form:"id" json:"id" xml:"id"`
	// 名前
	Name string `form:"name" json:"name" xml:"name"`
}

// Validate validates the UserTiny media type instance.
func (mt *UserTiny) Validate() (err error) {

	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	return
}

// UserCollection is the media type for an array of User (default view)
//
// Identifier: application/vnd.user+json; type=collection; view=default
type UserCollection []*User

// Validate validates the UserCollection media type instance.
func (mt UserCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// UserCollection is the media type for an array of User (tiny view)
//
// Identifier: application/vnd.user+json; type=collection; view=tiny
type UserTinyCollection []*UserTiny

// Validate validates the UserTinyCollection media type instance.
func (mt UserTinyCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// example (default view)
//
// Identifier: application/vnd.validation+json; view=default
type Validation struct {
	// デフォルト値
	DefaultType string `form:"defaultType" json:"defaultType" xml:"defaultType"`
	// メールアドレス
	Email string `form:"email" json:"email" xml:"email"`
	// 列挙型
	EnumType string `form:"enumType" json:"enumType" xml:"enumType"`
	// id
	ID int `form:"id" json:"id" xml:"id"`
	// 数字（1〜10）
	IntegerType int `form:"integerType" json:"integerType" xml:"integerType"`
	// デフォルト値
	Reg string `form:"reg" json:"reg" xml:"reg"`
	// 文字（1~10文字）
	StringType string `form:"stringType" json:"stringType" xml:"stringType"`
}

// Validate validates the Validation media type instance.
func (mt *Validation) Validate() (err error) {

	if mt.StringType == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "stringType"))
	}
	if mt.Email == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "email"))
	}
	if mt.EnumType == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "enumType"))
	}
	if mt.DefaultType == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "defaultType"))
	}
	if mt.Reg == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "reg"))
	}
	return
}
