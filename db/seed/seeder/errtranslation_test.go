package seeder

import (
	"airbnb-user-be/internal/pkg/env"
	"airbnb-user-be/internal/pkg/gorm"
	"testing"
)

func TestSeedErrTranslation(t *testing.T) {
	t.Log("seeding error translation...")
	env.InitEnv(envOps)
	config := env.ProvideEnv().DB
	engine := gorm.NewORM(gorm.Options{Config: config})
	if err := SeedErrTranslation(*engine.DB); err != nil {
		t.Error("failed to seed error translation", err)
	}
}
