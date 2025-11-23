package utility

import (
	"fmt"
	"os"
)

func GetString(key string) (string, error) {
	if val := os.Getenv(key); val != "" {
		return val, nil
	}

	return "", fmt.Errorf("MISSING ENV KEY: %s", key)
}
