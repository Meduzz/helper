package utilz

import "os"

// Env - dig out key out of environ, use defaultz if empty.
func Env(key, defaultz string) string {
	val := os.Getenv(key)

	if val == "" {
		val = defaultz
	}

	return val
}
