package env

import (
	"os"
)

func GetStage() Stage {
	stageEnv := os.Getenv("STAGE")
	return ParseString(stageEnv)
}
