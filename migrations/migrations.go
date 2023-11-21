package migrations

import (
	"go_pdf/internal/entities"
	"gorm.io/gorm"
)

var models = []interface{}{
	&entities.User{},
}

func AutoMigration(db *gorm.DB) {
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(models...)
}
