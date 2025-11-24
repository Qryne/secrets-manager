package utility

import (
	"fmt"
	"os"
)

func GetENVString(key string) (string, error) {
	val, ok := os.LookupEnv(key)

	if !ok {
		return "", fmt.Errorf("MISSING ENV KEY: %s", key)
	}

	return val, nil
}
