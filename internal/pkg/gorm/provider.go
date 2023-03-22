package gorm

import (
	"airbnb-user-be/internal/pkg/gorm/config"
	"context"

	"gorm.io/gorm"
)

const Instance string = "ORM"

type Engine struct {
	DB  *gorm.DB
	Ctx *context.Context
	config.Config
}

func NewORM(config config.Config) *Engine {
	ctx := context.Background()
	engine := Engine{
		Ctx:    &ctx,
		Config: config,
	}

	engine.InitConnection()
	return &engine
}
