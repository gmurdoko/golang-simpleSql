package utils

import "os"

func GetEnv(key, defaultValue string) string {
	if envVal, exists := os.LookupEnv(key); exists {
		return envVal
	}
	return defaultValue
}
