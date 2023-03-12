package seeder

import (
	"airbnb-user-be/internal/pkg/env"
	"airbnb-user-be/internal/pkg/gorm"
	"testing"
)

func TestSeedCurrencyTranslation(t *testing.T) {
	t.Log("seeding currency translation...")
	config := env.ProvideEnv(envConfig).DB
	engine := gorm.ProvideORM(config)
	if err := SeedCurrencyTranslation(*engine.DB); err != nil {
		t.Error("failed to seed currency translation", err)
	}
}
