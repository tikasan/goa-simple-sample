// Code generated by goagen v1.1.0-dirty, command line:
// $ goagen
// --design=github.com/tikasan/goa-simple-sample/design
// --out=$(GOPATH)/src/github.com/tikasan/goa-simple-sample
// --version=v1.1.0
//
// API "goa simple sample": Application Media Types
//
// The content of this file is auto-generated, DO NOT MODIFY

package app

import (
	"github.com/goadesign/goa"
)

// example (default view)
//
// Identifier: application/vnd.arraytype+json; view=default
type Arraytype struct {
	// ID
	ID int `form:"ID" json:"ID" xml:"ID"`
	// 値
	Value string `form:"value" json:"value" xml:"value"`
}

// Validate validates the Arraytype media type instance.
func (mt *Arraytype) Validate() (err error) {

	if mt.Value == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "value"))
	}
	return
}

// ArraytypeCollection is the media type for an array of Arraytype (default view)
//
// Identifier: application/vnd.arraytype+json; type=collection; view=default
type ArraytypeCollection []*Arraytype

// Validate validates the ArraytypeCollection media type instance.
func (mt ArraytypeCollection) Validate() (err error) {
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
// Identifier: application/vnd.integertype+json; view=default
type Integertype struct {
	// ID
	ID int `form:"ID" json:"ID" xml:"ID"`
}

// example (default view)
//
// Identifier: application/vnd.messagetype+json; view=default
type Messagetype struct {
	// メッセージ
	Message string `form:"message" json:"message" xml:"message"`
}

// Validate validates the Messagetype media type instance.
func (mt *Messagetype) Validate() (err error) {
	if mt.Message == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "message"))
	}
	return
}

// example (default view)
//
// Identifier: application/vnd.validationtype+json; view=default
type Validationtype struct {
	// ID
	ID int `form:"ID" json:"ID" xml:"ID"`
	// デフォルト値
	DefaultType string `form:"defaultType" json:"defaultType" xml:"defaultType"`
	// メールアドレス
	Email string `form:"email" json:"email" xml:"email"`
	// 列挙型
	EnumType string `form:"enumType" json:"enumType" xml:"enumType"`
	// 数字（1〜10）
	IntegerType int `form:"integerType" json:"integerType" xml:"integerType"`
	// デフォルト値
	Reg string `form:"reg" json:"reg" xml:"reg"`
	// 文字（1~10文字）
	StringType string `form:"stringType" json:"stringType" xml:"stringType"`
}

// Validate validates the Validationtype media type instance.
func (mt *Validationtype) Validate() (err error) {

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

// example (default view)
//
// Identifier: application/vnd.viewtype+json; view=default
type Viewtype struct {
	// ID
	ID int `form:"ID" json:"ID" xml:"ID"`
	// 値
	Value string `form:"value" json:"value" xml:"value"`
}

// Validate validates the Viewtype media type instance.
func (mt *Viewtype) Validate() (err error) {

	if mt.Value == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "value"))
	}
	return
}

// example (tiny view)
//
// Identifier: application/vnd.viewtype+json; view=tiny
type ViewtypeTiny struct {
	// ID
	ID int `form:"ID" json:"ID" xml:"ID"`
}
