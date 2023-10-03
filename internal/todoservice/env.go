package todoservice

import "blaze/pkg/util"

type EnvConfig struct {
	Environment string
}

var envVarMappings = util.EnvMapping{
	"ENVIRONMENT": &env.Environment,
}

var env = &EnvConfig{}

func InitEnv() *EnvConfig {
	for key, goVar := range envVarMappings {
		*goVar = util.GetEnvOrPanic(key)
	}

	return env
}
