package migration

import (
	currencymodule "airbnb-user-be/internal/app/currency"
	localemodule "airbnb-user-be/internal/app/locale"
	translationmodule "airbnb-user-be/internal/app/translation"
	usermodule "airbnb-user-be/internal/app/user"
	orm "airbnb-user-be/internal/pkg/gorm"
	"airbnb-user-be/internal/pkg/log"

	"gorm.io/gorm"
)

func MigrateUp(db gorm.DB) {
	models := []interface{}{
		&localemodule.Locale{},
		&translationmodule.ErrTranslation{},
		&currencymodule.Currency{},
		&currencymodule.CurrencyTranslation{},
		&usermodule.User{},
		&usermodule.UserDefaultSetting{},
	}

	if err := db.AutoMigrate(models...); err != nil {
		log.Fatal(orm.Instance, "failed to run migration", err)
	}
}
