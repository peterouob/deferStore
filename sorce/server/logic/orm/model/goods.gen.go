// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

const TableNameGood = "goods"

// Good mapped from table <goods>
type Good struct {
	ID            string          `gorm:"column:id;primaryKey" json:"id"`
	Name          string          `gorm:"column:name;not null" json:"name"`
	Img           string          `gorm:"column:img;not null" json:"img"`
	Price         decimal.Decimal `gorm:"column:price;not null" json:"price"`
	OriginalPrice decimal.Decimal `gorm:"column:original_price;not null" json:"original_price"`
	CreatedAt     time.Time       `gorm:"column:created_at" json:"created_at"`
	Desc          string          `gorm:"column:desc;not null" json:"desc"`
	UpdatedAt     time.Time       `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt     gorm.DeletedAt  `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName Good's table name
func (*Good) TableName() string {
	return TableNameGood
}
