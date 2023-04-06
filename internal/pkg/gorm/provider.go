package gorm

import (
	"airbnb-user-be/internal/pkg/env"
	"airbnb-user-be/internal/pkg/gorm/config"
	"airbnb-user-be/internal/pkg/log"
	"context"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const Instance string = "ORM"

type Options struct {
	config.Config
}

type Engine struct {
	Options
	DB  *gorm.DB
	Ctx *context.Context
}

func NewORM(options Options) *Engine {
	ctx := context.Background()
	engine := Engine{
		Ctx:     &ctx,
		Options: options,
	}

	engine.InitConnection()
	log.Event(Instance, fmt.Sprintf("connected to %s:%s", env.CONFIG.DB.Host, env.CONFIG.DB.Name))
	return &engine
}

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

	var config gorm.Config
	if env.CONFIG.Stage == string(env.StageLocal) {
		config.Logger = logger.Default.LogMode(logger.Info)
	}

	db, err := gorm.Open(postgres.Open(dsn), &config)
	if err != nil {
		log.Fatal(Instance, "failed to init db connection", err)
	}

	g.DB = db
}
