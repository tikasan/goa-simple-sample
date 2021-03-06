// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "goa simple sample": Models
//
// Command:
// $ goagen
// --design=github.com/pei0804/goa-simple-sample/design
// --out=$(GOPATH)/src/github.com/pei0804/goa-simple-sample
// --version=v1.1.0

package models

import (
	"context"
	"time"

	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
	"github.com/pei0804/goa-simple-sample/app"
)

// celler account
type Account struct {
	ID        int      `gorm:"primary_key"` // primary key
	Bottles   []Bottle // has many Bottles
	Email     string
	Name      string
	CreatedAt time.Time  // timestamp
	DeletedAt *time.Time // nullable timestamp (soft delete)
	UpdatedAt time.Time  // timestamp
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m Account) TableName() string {
	return "accounts"

}

// AccountDB is the implementation of the storage interface for
// Account.
type AccountDB struct {
	Db *gorm.DB
}

// NewAccountDB creates a new storage type.
func NewAccountDB(db *gorm.DB) *AccountDB {
	return &AccountDB{Db: db}
}

// DB returns the underlying database.
func (m *AccountDB) DB() interface{} {
	return m.Db
}

// AccountStorage represents the storage interface.
type AccountStorage interface {
	DB() interface{}
	List(ctx context.Context) ([]*Account, error)
	Get(ctx context.Context, id int) (*Account, error)
	Add(ctx context.Context, account *Account) error
	Update(ctx context.Context, account *Account) error
	Delete(ctx context.Context, id int) error

	ListAccount(ctx context.Context) []*app.Account
	OneAccount(ctx context.Context, id int) (*app.Account, error)
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m *AccountDB) TableName() string {
	return "accounts"

}

// CRUD Functions

// Get returns a single Account as a Database Model
// This is more for use internally, and probably not what you want in  your controllers
func (m *AccountDB) Get(ctx context.Context, id int) (*Account, error) {
	defer goa.MeasureSince([]string{"goa", "db", "account", "get"}, time.Now())

	var native Account
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&native).Error
	if err == gorm.ErrRecordNotFound {
		return nil, err
	}

	return &native, err
}

// List returns an array of Account
func (m *AccountDB) List(ctx context.Context) ([]*Account, error) {
	defer goa.MeasureSince([]string{"goa", "db", "account", "list"}, time.Now())

	var objs []*Account
	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return objs, nil
}

// Add creates a new record.
func (m *AccountDB) Add(ctx context.Context, model *Account) error {
	defer goa.MeasureSince([]string{"goa", "db", "account", "add"}, time.Now())

	err := m.Db.Create(model).Error
	if err != nil {
		goa.LogError(ctx, "error adding Account", "error", err.Error())
		return err
	}

	return nil
}

// Update modifies a single record.
func (m *AccountDB) Update(ctx context.Context, model *Account) error {
	defer goa.MeasureSince([]string{"goa", "db", "account", "update"}, time.Now())

	obj, err := m.Get(ctx, model.ID)
	if err != nil {
		goa.LogError(ctx, "error updating Account", "error", err.Error())
		return err
	}
	err = m.Db.Model(obj).Updates(model).Error

	return err
}

// Delete removes a single record.
func (m *AccountDB) Delete(ctx context.Context, id int) error {
	defer goa.MeasureSince([]string{"goa", "db", "account", "delete"}, time.Now())

	var obj Account

	err := m.Db.Delete(&obj, id).Error

	if err != nil {
		goa.LogError(ctx, "error deleting Account", "error", err.Error())
		return err
	}

	return nil
}
