package seeder

import (
	"airbnb-user-be/internal/pkg/env"
	"airbnb-user-be/internal/pkg/gorm"
	"testing"
)

func TestSeedLocale(t *testing.T) {
	t.Log("seeding locale...")
	env.InitEnv(envConfig)
	config := env.ProvideEnv().DB
	engine := gorm.NewORM(config)
	if err := SeedLocale(*engine.DB); err != nil {
		t.Error("failed to seed locale", err)
	}
}
