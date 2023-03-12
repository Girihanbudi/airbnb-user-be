package seeder

import (
	"airbnb-user-be/internal/pkg/env"
	"airbnb-user-be/internal/pkg/gorm"
	"testing"
)

func TestSeedCurrency(t *testing.T) {
	t.Log("seeding currency...")
	config := env.ProvideEnv(envConfig).DB
	engine := gorm.ProvideORM(config)
	if err := SeedCurrency(*engine.DB); err != nil {
		t.Error("failed to seed currency", err)
	}
}
