package util

import "os"

// GetEnvOrPanic returns the value of the environment variable with the given key.
func GetEnvOrPanic(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic(key + " environment variable is not set")
	}
	return value
}

// GetEnvOrDefault returns the value of the environment variable with the given key,
// or a default value if the environment variable is not set.
func GetEnvOrDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
