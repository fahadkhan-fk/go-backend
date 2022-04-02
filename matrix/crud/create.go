package crud

import (
	"github.com/jinzhu/gorm"
)

// Create creates the record
func Create(db *gorm.DB, model interface{}) *gorm.DB {
	return db.Create(model)
}
