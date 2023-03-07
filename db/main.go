package main

import (
	m "airbnb-user-be/db/migration"
	"airbnb-user-be/internal/pkg/env"
	"airbnb-user-be/internal/pkg/gorm"
	"airbnb-user-be/internal/pkg/log"
	"flag"
)

var Instance = "Migration"

func main() {
	cmd := flag.String("migration", "", "migration direction (up/down)")
	flag.Parse()
	if cmd == nil || *cmd == "" {
		log.Fatal(Instance, "migration failed with arguments not found", nil)
	}

	config := env.ProvideEnv().DB
	engine := gorm.ProvideORM(config)

	switch *cmd {
	case m.MigrationUp.String():
		log.Event(Instance, "migrating...")
		m.MigrateUp(*engine.DB)
		log.Event(Instance, "finish migrating")
	default:
		log.Fatal(Instance, "unknown command", nil)
	}
}
