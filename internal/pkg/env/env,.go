package env

import "os"

var (
	env string
)

const (
	EnvLocal = "local"
	EnvDev   = "dev"
	EnvProd  = "prod"
)

func init() {
	env = os.Getenv("ENV")
	if env == "" {
		env = EnvLocal
	}
}

func GetEnv() string {
	return env
}
