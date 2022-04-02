package crud

import (
	"github.com/jinzhu/gorm"
)

// Update updates the record
func Update(db *gorm.DB, model interface{}) *gorm.DB {
	return db.Save(model)
}
