package config

import (
	"os"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func DbConfig() *gorm.Config {
	env := os.Getenv("ENV")

	if env == "prod" {
		return &gorm.Config{
			Logger:                                   logger.Discard,
			PrepareStmt:                              false,
			SkipDefaultTransaction:                   true,
			DisableAutomaticPing:                     true,
			DisableForeignKeyConstraintWhenMigrating: true,
		}
	}

	// dev
	return &gorm.Config{
		Logger:                                   logger.Default.LogMode(logger.Info),
		PrepareStmt:                              true,
		SkipDefaultTransaction:                   false,
		DisableAutomaticPing:                     false,
		DisableForeignKeyConstraintWhenMigrating: false,
	}
}
