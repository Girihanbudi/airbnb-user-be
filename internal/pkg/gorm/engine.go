package gorm

import (
	"fmt"

	"airbnb-user-be/internal/pkg/log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func (g *Engine) InitConnection() {
	dsn := fmt.Sprintf(`host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s`,
		g.Host,
		g.User,
		g.Password,
		g.Name,
		g.Port,
		g.SslMode,
		g.Timezone,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(Instance, "failed to init db connection", err)
	}

	g.DB = db
}
