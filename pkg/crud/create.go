package crud

import (
	"gorm.io/gorm"
)

// Create creates the record
func Create(db *gorm.DB, model interface{}) *gorm.DB {
	return db.Create(model)
}
