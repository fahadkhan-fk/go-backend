package crud

import (
	"github.com/jinzhu/gorm"
)

// Get returns a single record
func Get(db *gorm.DB, model interface{}, id string) *gorm.DB {
	return db.Where("id = ?", id).Find(model)
}
