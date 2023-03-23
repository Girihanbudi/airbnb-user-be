package seeder

import (
	"airbnb-user-be/internal/pkg/env"
	"airbnb-user-be/internal/pkg/gorm"
	"testing"
)

func TestSeedCurrency(t *testing.T) {
	t.Log("seeding currency...")
	env.InitEnv(envConfig)
	config := env.ProvideEnv().DB
	engine := gorm.NewORM(config)
	if err := SeedCurrency(*engine.DB); err != nil {
		t.Error("failed to seed currency", err)
	}
}
