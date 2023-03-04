package gorm

import (
	"context"

	"gorm.io/gorm"
)

var Instance string = "ORM"

type Engine struct {
	DB  *gorm.DB
	Ctx *context.Context
	Config
}

func ProvideORM(config Config) *Engine {
	ctx := context.Background()
	return &Engine{
		Ctx:    &ctx,
		Config: config,
	}
}
