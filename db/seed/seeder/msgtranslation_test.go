package seeder

import (
	"airbnb-user-be/internal/pkg/env"
	"airbnb-user-be/internal/pkg/gorm"
	"testing"
)

func TestSeedMsgTranslation(t *testing.T) {
	t.Log("seeding msg translation...")
	env.InitEnv(envOps)
	config := env.ProvideEnv().DB
	engine := gorm.NewORM(gorm.Options{Config: config})
	if err := SeedMsgTranslation(*engine.DB); err != nil {
		t.Error("failed to seed msg translation", err)
	}
}
