package seeder

import (
	"airbnb-user-be/internal/pkg/env"
	"airbnb-user-be/internal/pkg/gorm"
	"testing"
)

func TestSeedLocale(t *testing.T) {
	t.Log("seeding locale...")
	config := env.ProvideEnv(envConfig).DB
	engine := gorm.ProvideORM(config)
	if err := SeedLocale(*engine.DB); err != nil {
		t.Error("failed to seed locale", err)
	}
}
