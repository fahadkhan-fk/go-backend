package crud

import (
	"gorm.io/gorm"
)

// Delete deletes the record
func Delete(db *gorm.DB, model interface{}) *gorm.DB {
	return db.Unscoped().Delete(model)
}
