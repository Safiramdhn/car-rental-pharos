package helpers

import "gorm.io/gorm"

func Paginate(page uint, limit uint) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if limit > uint(0) {
			offset := (page - 1) * limit
			return db.Offset(int(offset)).Limit(int(limit))
		}
		return db
	}
}
