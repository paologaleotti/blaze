package api

import "blaze/pkg/util"

type EnvConfig struct {
	Environment string
	DatabaseUrl string
}

func InitEnv() EnvConfig {
	return EnvConfig{
		Environment: util.GetEnvOrDefault("ENVIRONMENT", "dev"),
		DatabaseUrl: util.GetEnvOrDefault("DATABASE_URL", "db.sqlite3"),
	}
}
