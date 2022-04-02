package crud

import (
	"github.com/jinzhu/gorm"
)

// Index returns all records
func Index(db *gorm.DB, model interface{}) *gorm.DB {
	return db.Find(model)
}
