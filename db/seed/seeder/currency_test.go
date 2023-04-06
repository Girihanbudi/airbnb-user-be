package seeder

import (
	"airbnb-user-be/internal/pkg/env"
	"airbnb-user-be/internal/pkg/gorm"
	"testing"
)

func TestSeedCurrency(t *testing.T) {
	t.Log("seeding currency...")
	env.InitEnv(envOps)
	config := env.ProvideEnv().DB
	engine := gorm.NewORM(gorm.Options{Config: config})
	if err := SeedCurrency(*engine.DB); err != nil {
		t.Error("failed to seed currency", err)
	}
}
