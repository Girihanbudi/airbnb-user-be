package migration

import (
	"airbnb-user-be/internal/app/locale"
	orm "airbnb-user-be/internal/pkg/gorm"
	"airbnb-user-be/internal/pkg/log"

	"gorm.io/gorm"
)

func MigrateUp(db gorm.DB) {
	models := []interface{}{
		&locale.Locale{},
	}

	if err := db.AutoMigrate(models...); err != nil {
		log.Fatal(orm.Instance, "failed to run migration", err)
	}
}
