package seeder

import (
	"airbnb-user-be/internal/pkg/env"
	"testing"

	"airbnb-user-be/internal/pkg/gorm"
)

func TestSeedAll(t *testing.T) {
	env.InitEnv(envOps)
	config := env.ProvideEnv().DB
	engine := gorm.NewORM(gorm.Options{Config: config})
	db := *engine.DB

	t.Log("seeding error translation...")
	if err := SeedErrTranslation(db); err != nil {
		t.Error("failed to seed error translation", err)
	}

	t.Log("seeding locale...")
	if err := SeedLocale(db); err != nil {
		t.Error("failed to seed locale", err)
	}

	t.Log("seeding currency...")
	if err := SeedCurrency(db); err != nil {
		t.Error("failed to seed currency", err)
	}

	t.Log("seeding currency translation...")
	if err := SeedCurrencyTranslation(db); err != nil {
		t.Error("failed to seed currency translation", err)
	}

	t.Log("finish seeding")
}
