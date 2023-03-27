package seeder

import (
	"airbnb-user-be/internal/pkg/env"
	"airbnb-user-be/internal/pkg/gorm"
	"testing"
)

func TestSeedCountry(t *testing.T) {
	t.Log("seeding country...")
	env.InitEnv(envConfig)
	config := env.ProvideEnv().DB
	engine := gorm.NewORM(config)
	if err := SeedCountry(*engine.DB); err != nil {
		t.Error("failed to seed country", err)
	}
}
