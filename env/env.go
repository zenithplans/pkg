// Provides basic methods to interact with environment variables.
package env

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Get returns the environment variable value as a string,
// or the fallback if the variable is not set.
func Get(key string, fallback string) string {
	val, exists := os.LookupEnv(key)

	if !exists || val == "" {
		return fallback
	}

	return strings.TrimSpace(val)
}

// GetInt returns the environment variable value parsed as an
// integer, or the fallback if the variable is not set or
// invalid.
func GetInt(key string, fallback int) int {

	valStr, exists := os.LookupEnv(key)
	if !exists || valStr == "" {
		return fallback
	}

	valInt, err := strconv.Atoi(valStr)
	if err != nil {
		panicErr := fmt.Errorf(
			"environment variable [%s] is not a valid integer: %w",
			key,
			err,
		)
		panic(panicErr)
	}

	return valInt
}

// MustGet returns the environment variable as a string,
// or panics if it is not set.
func MustGet(key string) string {

	val, exists := os.LookupEnv(key)
	if !exists {
		err := fmt.Errorf(
			"missing environment variable: %s",
			key,
		)

		panic(err)
	}

	return strings.TrimSpace(val)
}

// MustGetInt returns the environment variable parsed as an
// integer, or panics if it is not set or not a valid integer.
func MustGetInt(key string) int {
	valStr := MustGet(key)
	valInt, err := strconv.Atoi(valStr)
	if err != nil {
		panicErr := fmt.Errorf(
			"environment variable [%s] must be an integer: %w",
			key,
			err,
		)
		panic(panicErr)
	}

	return valInt
}
