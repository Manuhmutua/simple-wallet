package postgres

import (
	"wallet/storage"
	"wallet/user"
)

// Migrate updates the db with new columns, and tables
func Migrate(database *storage.Database) {
	database.DB.AutoMigrate(user.User{})
}