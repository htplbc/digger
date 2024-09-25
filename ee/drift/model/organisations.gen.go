// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameOrganisation = "organisations"

// Organisation mapped from table <organisations>
type Organisation struct {
	ID             string         `gorm:"column:id;primaryKey;default:uuid_generate_v4()" json:"id"`
	CreatedAt      time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt      time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	Name           string         `gorm:"column:name" json:"name"`
	ExternalSource string         `gorm:"column:external_source" json:"external_source"`
	ExternalID     string         `gorm:"column:external_id" json:"external_id"`
}

// TableName Organisation's table name
func (*Organisation) TableName() string {
	return TableNameOrganisation
}
