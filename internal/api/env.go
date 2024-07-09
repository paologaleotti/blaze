package api

import "blaze/pkg/util"

type EnvConfig struct {
	Environment string
}

func InitEnv() EnvConfig {
	return EnvConfig{
		Environment: util.GetEnvOrDefault("ENVIRONMENT", "dev"),
	}
}
