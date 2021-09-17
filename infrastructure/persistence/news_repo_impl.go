package persistence

import "gorm.io/gorm"

type NewsRepoImpl struct {
	db *gorm.DB
}
