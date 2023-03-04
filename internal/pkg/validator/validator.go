package validator

import (
	"sync"

	goValidator "github.com/go-playground/validator/v10"
)

var singleValidator sync.Once

var (
	validator *goValidator.Validate
)

func InitValidator() {
	// do only once
	singleValidator.Do(func() {
		validator = goValidator.New()
	})
}

func GetValidator() *goValidator.Validate {
	return validator
}
