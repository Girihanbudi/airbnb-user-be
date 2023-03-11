package seeder

import (
	"airbnb-user-be/internal/pkg/env"
	"airbnb-user-be/internal/pkg/gorm"
	"testing"
)

func TestSeedErrTranslation(t *testing.T) {
	t.Log("seeding error translation...")
	config := env.ProvideEnv(envConfig).DB
	engine := gorm.ProvideORM(config)
	if err := SeedErrTranslation(*engine.DB); err != nil {
		t.Error("failed to seed error translation", err)
	}
}
