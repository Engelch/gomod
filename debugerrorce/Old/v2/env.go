package debugerrorce

import (
	"errors"
	"log"
	"os"
	"strconv"
	"strings"
)

// GetEnvValue retrieves a value from the ENV if exists, otherwise
// returns an error
func GetEnvValue(key string) (string, error) {
	val, present := os.LookupEnv(key)
	if !present {
		return "", errors.New(key + " env variable must be defined")
	}
	return val, nil
}

// GetEnvValueOrDefaultBool returns the ENV boolean value for the given key or the default one
func GetEnvValueOrDefaultBool(key string, defaultValue bool) bool {
	val, err := GetEnvValue(key)
	if err != nil {
		return defaultValue
	}
	return strings.ToLower(val) == "true"
}

// GetEnvValueOrDefaultString returns the ENV string value for the given key or the default one
func GetEnvValueOrDefaultString(key string, defaultValue string) string {
	val, err := GetEnvValue(key)
	if err != nil {
		return defaultValue
	}
	return val
}

// GetEnvValueOrDefaultInt returns the ENV int value for the given key or the default one if none could
// be found or if there is an error parsing the string to an int
func GetEnvValueOrDefaultInt(key string, defaultValue int) int {
	val, err := GetEnvValue(key)
	if err != nil {
		return defaultValue
	}
	i, err := strconv.Atoi(val)
	if err != nil {
		return defaultValue
	}
	return i
}

// FatalGetEnvValue does the same as getEnvValue, but exits with an error
func FatalGetEnvValue(key string) string {
	val, err := GetEnvValue(key)
	if err != nil {
		log.Fatalf("%v", err)
	}
	return val
}

// eof
