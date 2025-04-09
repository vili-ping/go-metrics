package utils

import "os"

func IsEnvSet(key string) bool {
	_, exists := os.LookupEnv(key)
	return exists
}
