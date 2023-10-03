package util

import "os"

type EnvMapping map[string]*string

func GetEnvOrPanic(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic(key + " environment variable is not set")
	}
	return value
}
