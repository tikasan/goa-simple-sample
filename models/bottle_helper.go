// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "goa simple sample": Model Helpers
//
// Command:
// $ goagen
// --design=github.com/tikasan/goa-simple-sample/design
// --out=$(GOPATH)/src/github.com/tikasan/goa-simple-sample
// --version=v1.1.0

package models

import (
	"context"
	"time"

	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
	"github.com/tikasan/goa-simple-sample/app"
)

// MediaType Retrieval Functions

// ListBottle returns an array of view: default.
func (m *BottleDB) ListBottle(ctx context.Context, accountID int) []*app.Bottle {
	defer goa.MeasureSince([]string{"goa", "db", "bottle", "listbottle"}, time.Now())

	var native []*Bottle
	var objs []*app.Bottle
	err := m.Db.Scopes(BottleFilterByAccount(accountID, m.Db)).Table(m.TableName()).Find(&native).Error

	if err != nil {
		goa.LogError(ctx, "error listing Bottle", "error", err.Error())
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.BottleToBottle())
	}

	return objs
}

// BottleToBottle loads a Bottle and builds the default view of media type Bottle.
func (m *Bottle) BottleToBottle() *app.Bottle {
	bottle := &app.Bottle{}
	tmp1 := &m.Account
	bottle.Account = tmp1.AccountToAccount() // %!s(MISSING)
	bottle.ID = m.ID
	bottle.Name = m.Name
	bottle.Quantity = m.Quantity

	return bottle
}

// OneBottle loads a Bottle and builds the default view of media type Bottle.
func (m *BottleDB) OneBottle(ctx context.Context, id int, accountID int) (*app.Bottle, error) {
	defer goa.MeasureSince([]string{"goa", "db", "bottle", "onebottle"}, time.Now())

	var native Bottle
	err := m.Db.Scopes(BottleFilterByAccount(accountID, m.Db)).Table(m.TableName()).Preload("BottleCategories").Preload("Account").Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error getting Bottle", "error", err.Error())
		return nil, err
	}

	view := *native.BottleToBottle()
	return &view, err
}

// MediaType Retrieval Functions

// ListBottleRelation returns an array of view: relation.
func (m *BottleDB) ListBottleRelation(ctx context.Context, accountID int) []*app.BottleRelation {
	defer goa.MeasureSince([]string{"goa", "db", "bottle", "listbottlerelation"}, time.Now())

	var native []*Bottle
	var objs []*app.BottleRelation
	err := m.Db.Scopes(BottleFilterByAccount(accountID, m.Db)).Table(m.TableName()).Find(&native).Error

	if err != nil {
		goa.LogError(ctx, "error listing Bottle", "error", err.Error())
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.BottleToBottleRelation())
	}

	return objs
}

// BottleToBottleRelation loads a Bottle and builds the relation view of media type Bottle.
func (m *Bottle) BottleToBottleRelation() *app.BottleRelation {
	bottle := &app.BottleRelation{}
	tmp1 := &m.Account
	bottle.Account = tmp1.AccountToAccount() // %!s(MISSING)
	for i3 := range m.Categories {
		tmp3 := &m.Categories[i3]
		bottle.Categories = append(bottle.Categories, tmp3.CategoryToCategory())
	}
	bottle.ID = m.ID
	bottle.Name = m.Name
	bottle.Quantity = m.Quantity

	return bottle
}

// OneBottleRelation loads a Bottle and builds the relation view of media type Bottle.
func (m *BottleDB) OneBottleRelation(ctx context.Context, id int, accountID int) (*app.BottleRelation, error) {
	defer goa.MeasureSince([]string{"goa", "db", "bottle", "onebottlerelation"}, time.Now())

	var native Bottle
	err := m.Db.Scopes(BottleFilterByAccount(accountID, m.Db)).Table(m.TableName()).Preload("BottleCategories").Preload("Account").Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error getting Bottle", "error", err.Error())
		return nil, err
	}

	view := *native.BottleToBottleRelation()
	return &view, err
}
