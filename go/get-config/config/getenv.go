package config

import (
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

func init() {
	_, filename, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(filename)
	filepath := path.Join(basepath, "../../.env")

	// In case we are testing, try to load .env using relative path and ignore the error.
	err := godotenv.Load(filepath)
	if err != nil && !strings.HasSuffix(err.Error(), "no such file or directory") {
		panic(err)
	}
}

func TryGet(key string) (string, bool) {
	var value string
	value = strings.TrimSpace(os.Getenv(key))
	if value == "" {
		value = strings.ToUpper(strings.TrimSpace(os.Getenv(key)))

		if value == "" {
			return "", false
		}
	}

	return value, true
}

func TryGetInt(key string) (int, bool) {
	value, ok := TryGet(key)
	if !ok {
		return 0, false
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}

	return intValue, true
}

func TryGetBool(key string) (bool, bool) {
	value, ok := TryGet(key)
	if !ok {
		return false, false
	}

	bValue, err := strconv.ParseBool(value)
	if err != nil {
		panic(err)
	}

	return bValue, true
}
