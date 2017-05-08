// Code generated by goagen v1.1.0-dirty, command line:
// $ goagen
// --design=github.com/tikasan/goa-simple-sample/design
// --out=$(GOPATH)/src/github.com/tikasan/goa-simple-sample
// --version=v1.1.0
//
// API "goa simple sample": Models
//
// The content of this file is auto-generated, DO NOT MODIFY

package models

import (
	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
	"golang.org/x/net/context"
	"time"
)

// celler bottle category
type BottleCategory struct {
	ID         int        `gorm:"primary_key"` // primary key
	BottleID   int        // Belongs To Bottle
	CategoryID int        // has many BottleCategory
	CreatedAt  time.Time  // timestamp
	DeletedAt  *time.Time // nullable timestamp (soft delete)
	UpdatedAt  time.Time  // timestamp
	Bottle     Bottle
	Category   Category
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m BottleCategory) TableName() string {
	return "bottle_categories"

}

// BottleCategoryDB is the implementation of the storage interface for
// BottleCategory.
type BottleCategoryDB struct {
	Db *gorm.DB
}

// NewBottleCategoryDB creates a new storage type.
func NewBottleCategoryDB(db *gorm.DB) *BottleCategoryDB {
	return &BottleCategoryDB{Db: db}
}

// DB returns the underlying database.
func (m *BottleCategoryDB) DB() interface{} {
	return m.Db
}

// BottleCategoryStorage represents the storage interface.
type BottleCategoryStorage interface {
	DB() interface{}
	List(ctx context.Context) ([]*BottleCategory, error)
	Get(ctx context.Context, id int) (*BottleCategory, error)
	Add(ctx context.Context, bottlecategory *BottleCategory) error
	Update(ctx context.Context, bottlecategory *BottleCategory) error
	Delete(ctx context.Context, id int) error
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m *BottleCategoryDB) TableName() string {
	return "bottle_categories"

}

// Belongs To Relationships

// BottleCategoryFilterByBottle is a gorm filter for a Belongs To relationship.
func BottleCategoryFilterByBottle(bottleID int, originaldb *gorm.DB) func(db *gorm.DB) *gorm.DB {

	if bottleID > 0 {

		return func(db *gorm.DB) *gorm.DB {
			return db.Where("bottle_id = ?", bottleID)

		}
	}
	return func(db *gorm.DB) *gorm.DB { return db }
}

// Belongs To Relationships

// BottleCategoryFilterByCategory is a gorm filter for a Belongs To relationship.
func BottleCategoryFilterByCategory(categoryID int, originaldb *gorm.DB) func(db *gorm.DB) *gorm.DB {

	if categoryID > 0 {

		return func(db *gorm.DB) *gorm.DB {
			return db.Where("category_id = ?", categoryID)

		}
	}
	return func(db *gorm.DB) *gorm.DB { return db }
}

// CRUD Functions

// Get returns a single BottleCategory as a Database Model
// This is more for use internally, and probably not what you want in  your controllers
func (m *BottleCategoryDB) Get(ctx context.Context, id int) (*BottleCategory, error) {
	defer goa.MeasureSince([]string{"goa", "db", "bottleCategory", "get"}, time.Now())

	var native BottleCategory
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&native).Error
	if err == gorm.ErrRecordNotFound {
		return nil, err
	}

	return &native, err
}

// List returns an array of BottleCategory
func (m *BottleCategoryDB) List(ctx context.Context) ([]*BottleCategory, error) {
	defer goa.MeasureSince([]string{"goa", "db", "bottleCategory", "list"}, time.Now())

	var objs []*BottleCategory
	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return objs, nil
}

// Add creates a new record.
func (m *BottleCategoryDB) Add(ctx context.Context, model *BottleCategory) error {
	defer goa.MeasureSince([]string{"goa", "db", "bottleCategory", "add"}, time.Now())

	err := m.Db.Create(model).Error
	if err != nil {
		goa.LogError(ctx, "error adding BottleCategory", "error", err.Error())
		return err
	}

	return nil
}

// Update modifies a single record.
func (m *BottleCategoryDB) Update(ctx context.Context, model *BottleCategory) error {
	defer goa.MeasureSince([]string{"goa", "db", "bottleCategory", "update"}, time.Now())

	obj, err := m.Get(ctx, model.ID)
	if err != nil {
		goa.LogError(ctx, "error updating BottleCategory", "error", err.Error())
		return err
	}
	err = m.Db.Model(obj).Updates(model).Error

	return err
}

// Delete removes a single record.
func (m *BottleCategoryDB) Delete(ctx context.Context, id int) error {
	defer goa.MeasureSince([]string{"goa", "db", "bottleCategory", "delete"}, time.Now())

	var obj BottleCategory

	err := m.Db.Delete(&obj, id).Error

	if err != nil {
		goa.LogError(ctx, "error deleting BottleCategory", "error", err.Error())
		return err
	}

	return nil
}